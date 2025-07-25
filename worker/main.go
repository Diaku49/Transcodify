package main

import (
	"fmt"
	"log"

	"s3client"

	"github.com/Diaku49/FoodOrderSystem/worker/Redis"
	"github.com/Diaku49/FoodOrderSystem/worker/db"
	rmqclient "github.com/Diaku49/FoodOrderSystem/worker/rmq-worker"
)

func main() {
	db := db.Connect()
	err := s3client.InitS3Client()
	if err != nil {
		log.Fatalf("Storage failed: %v", err)
	}
	rmqClient := rmqclient.InitMQClient(db)
	redisClient := Redis.NewRedisClient()

	fmt.Println("Worker started")

	rmqClient.HandleUploadVideos(redisClient)
}
