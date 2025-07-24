package main

import (
	"log"
	"net/http"
	"s3client"

	"github.com/Diaku49/FoodOrderSystem/backend/internals/db"
	"github.com/Diaku49/FoodOrderSystem/backend/internals/email"
	"github.com/Diaku49/FoodOrderSystem/backend/internals/router"
	"github.com/Diaku49/FoodOrderSystem/backend/mq"
)

func main() {
	err := s3client.InitS3Client()
	if err != nil {
		log.Fatalf("Initializing s3client failed: %v", err)
	}

	database := db.Connect()
	if err := db.Migrate(database); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	err = email.InitMail()
	if err != nil {
		log.Fatalf("Initializing mail client failed: %v", err)
	}

	mqClient := mq.InitRabbitmqClient()
	r := router.SetupRouter(database, mqClient)

	log.Println("Server running on port:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
