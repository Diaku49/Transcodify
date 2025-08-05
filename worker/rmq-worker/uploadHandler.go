package rmqworker

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"s3client"
	"sync"

	"github.com/Diaku49/FoodOrderSystem/worker/Redis"
	"github.com/Diaku49/FoodOrderSystem/worker/db"
	"github.com/Diaku49/FoodOrderSystem/worker/model"
	amqp "github.com/rabbitmq/amqp091-go"
)

var res = map[string]string{
	"360p":  "480x360",
	"480p":  "720:480",
	"720p":  "1280:720",
	"1080p": "1920:1080",
}

func (mqc *MQClient) HandleUploadVideos(rdbc *Redis.RedisClient) error {

	messages, err := mqc.Channel.Consume(
		mqc.UploadVideoQueue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("couldnt get messages from broker: %w", err)
	}

	for msg := range messages {
		go func(msg amqp.Delivery) {

			var meta model.UploadedTempMetadata
			marshalError := json.Unmarshal(msg.Body, &meta)
			if marshalError != nil {
				log.Printf("Failed to unmarshal metadata: %v", marshalError)
				msg.Nack(false, false)
				return
			}

			// update redis status
			rdbc.UpdateVideoProgress(meta.Id, "10%% complete: Transcoding video")
			//transcode video
			paths, err := transcode(meta)
			if err != nil {
				log.Printf("Transcoding failed: %v", err)
				rdbc.DeleteMetadata(meta.Id) // redis
				msg.Nack(false, false)
				return
			}

			// upload and collect the url path
			var wg sync.WaitGroup
			videoVariantCh := make(chan model.VideoVariant, len(paths))

			for resolution, localPath := range paths {
				wg.Add(1)
				go func(resolution, path string) {
					defer wg.Done()
					key, err := s3client.S3Client.Upload(meta.Id, resolution, path)
					if err != nil {
						log.Printf("Upload failed for %s: %v", resolution, err)
						rdbc.DeleteMetadata(meta.Id) // redis
						return
					}
					videoVariantCh <- model.VideoVariant{
						Resolution: resolution,
						Key:        key,
					}
				}(resolution, localPath)
			}

			rdbc.UpdateVideoProgress(meta.Id, "50%% complete: Uploading video")
			wg.Wait()
			close(videoVariantCh)

			// final Video map
			var videos []model.VideoVariant
			for res := range videoVariantCh {
				videos = append(videos, res)
			}

			// Check which uploads failed so we delete them in system or put them in another queue

			// making the metadata records in db
			if err = db.CreateVideo(mqc.DB, meta.VideoName, videos); err != nil {
				log.Printf("DB insert failed: %v", err)
				rdbc.DeleteMetadata(meta.Id) // redis
				s3client.S3Client.Delete(meta.Id)
				return
			}

			// delete the temp metadata and tmp Transcoded videos
			rdbc.UpdateVideoProgress(meta.Id, "90%% complete: Cleanup")
			os.Remove(meta.Path)
			for _, path := range paths {
				os.Remove(path)
			}
			s3client.S3Client.Delete(meta.Id)

			msg.Ack(false)

			rdbc.UpdateVideoProgress(meta.Id, "Finished")
		}(msg)
	}

	return nil
}

func transcode(meta model.UploadedTempMetadata) (map[string]string, error) {
	// Create transcoded-videos directory if it doesn't exist
	if err := os.MkdirAll("../tmp/transcoded-videos", 0755); err != nil {
		return nil, fmt.Errorf("failed to create transcoded-videos directory: %w", err)
	}

	var wg sync.WaitGroup
	resultPath := make(map[string]string)
	var mu sync.Mutex

	for _, resolution := range meta.Resolutions {
		// setup goroutines
		size, ok := res[resolution]
		if !ok {
			log.Printf("Invalid resolution: %s", resolution)
			continue
		}
		wg.Add(1)

		// start transcoding
		go func(resolutionLabel, size string, resultpath map[string]string) {
			defer wg.Done()

			outputPath := fmt.Sprintf("../tmp/transcoded-videos/%s_%s_%s.mp4", meta.Id, meta.VideoName, resolutionLabel)

			cmd := exec.Command(
				"ffmpeg",
				"-i", meta.Path,
				"-vf", fmt.Sprintf("scale=%s", size),
				"-c:a", "copy",
				outputPath,
			)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				log.Printf("failed to transcode to %s: %v", resolutionLabel, err)
			}
			mu.Lock()
			resultpath[resolutionLabel] = outputPath
			mu.Unlock()

			log.Printf("Transcoded %s: %s", resolutionLabel, outputPath)

		}(resolution, size, resultPath)
	}

	wg.Wait()

	return resultPath, nil
}
