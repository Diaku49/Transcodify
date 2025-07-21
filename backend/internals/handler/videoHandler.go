package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/Diaku49/FoodOrderSystem/backend/Redis"
	"github.com/Diaku49/FoodOrderSystem/backend/internals/model"
	"github.com/Diaku49/FoodOrderSystem/backend/internals/repository"
	"github.com/Diaku49/FoodOrderSystem/backend/mq"
	util "github.com/Diaku49/FoodOrderSystem/backend/utilities"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type VideoHandler struct {
	VideoRepository *repository.VideoRepository
	RmqClient       *mq.MQClient
	RedisClient     *Redis.RedisClient
}

func NewVH(db *gorm.DB, mqc *mq.MQClient, rdbc *Redis.RedisClient) *VideoHandler {
	return &VideoHandler{
		VideoRepository: &repository.VideoRepository{DB: db},
		RmqClient:       mqc,
		RedisClient:     rdbc,
	}
}

func (vh *VideoHandler) GetAllVideos(w http.ResponseWriter, r *http.Request) {
	page, limit := util.GetPageLimit(r)
	offset := ((page - 1) * limit)

	videos, err := vh.VideoRepository.GetAllVideos(limit, offset)
	if err != nil {
		util.WriteJsonError(w, "db error", http.StatusInternalServerError, err)
		return
	}

	resp := transformVideoToResp(videos)

	util.WriteJsonSuccess(w, http.StatusOK, resp)
}

func (vh *VideoHandler) UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, (100 * 1024 * 1024))

	// Parsing multipart form
	err := r.ParseMultipartForm(100 << 20)
	if err != nil {
		util.WriteJsonError(w, "File too big or badly formatted", http.StatusBadRequest, err)
		return
	}

	// getting our file and metadata
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		util.WriteJsonError(w, "Missing file", http.StatusBadRequest, err)
		return
	}
	metadataStr := r.FormValue("metadata")
	var metadata model.UploadMetadata
	err = json.Unmarshal([]byte(metadataStr), &metadata)
	if err != nil {
		util.WriteJsonError(w, "Invalid metadata JSON", http.StatusBadRequest, err)
		return
	}

	// making uuid, path and saving file
	id := uuid.New().String()
	path := "tmp/uploads/" + id + "_" + fileHeader.Filename
	dst, err := os.Create(path)
	if err != nil {
		util.WriteJsonError(w, "Failed to save file", http.StatusInternalServerError, err)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		util.WriteJsonError(w, "Failed to write file", http.StatusInternalServerError, err)
		return
	}

	// Saving metadata in json file
	tempMeta := model.UploadedTempMetadata{
		Id:          id,
		VideoName:   metadata.VideoName,
		Resolutions: metadata.Resolutions,
		Path:        path,
	}

	err = vh.RedisClient.SaveMetadata(id, &tempMeta)
	if err != nil {
		util.WriteJsonError(w, "Failed to save metadata", http.StatusInternalServerError, err)
		os.Remove(tempMeta.Path)
	}

	// publishing the metadata to our worker
	msgBytes, err := json.Marshal(tempMeta)
	if err != nil {
		util.WriteJsonError(w, "Failed to marshal message", http.StatusInternalServerError, err)
		return
	}
	err = vh.RmqClient.Channel.Publish(
		vh.RmqClient.ExchangeName,
		vh.RmqClient.RoutingKey,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        msgBytes,
		},
	)
	if err != nil {
		util.WriteJsonError(w, "Failed publish message", http.StatusInternalServerError, err)
		return
	}

	resp := model.UploadVideoResp{
		ID:      id,
		Message: "Upload successfully",
	}

	util.WriteJsonSuccess(w, http.StatusOK, resp)
}

func (vh *VideoHandler) GetVideoInfoHandler(w http.ResponseWriter, r *http.Request) {
	videoId := chi.URLParam(r, "videoId")

	progress, err := vh.RedisClient.GetVideoProgress(videoId)
	if err != nil {
		util.WriteJsonError(w, "Action failed", http.StatusInternalServerError, err)
		return
	}

	resp := map[string]string{
		"status": progress,
	}

	util.WriteJsonSuccess(w, http.StatusOK, resp)
}

// func (vh *VideoHandler) Download(w http.ResponseWriter, r *http.Request) {
// 	videoId := chi.URLParam(r, "videoId")

// }

// func (vh *VideoHandler) Delete(w http.ResponseWriter, r *http.Request) {
// 	videoId := chi.URLParam(r, "videoId")

// }

func transformVideoToResp(videos []model.Video) *model.GetAllVideosResp {
	var videoResp []model.VideoResp
	if len(videoResp) != 0 {
		for _, v := range videos {
			variants := make([]model.VideoVariantResp, len(v.VideoVariant))
			if len(v.VideoVariant) != 0 {
				for i, vv := range v.VideoVariant {
					variants[i] = model.VideoVariantResp{
						ID:         vv.ID,
						Resolution: vv.Resolution,
						URL:        vv.URL,
					}
				}
			}
			videoResp = append(videoResp, model.VideoResp{
				ID:           v.ID,
				Name:         v.Name,
				VideoVariant: variants,
			})
		}
	}

	resp := model.GetAllVideosResp{
		Videos:  videoResp,
		Message: "videos fetched successfully",
	}

	return &resp
}
