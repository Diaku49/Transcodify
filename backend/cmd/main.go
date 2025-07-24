package main

import (
	"log"
	"net/http"
	"os"
	"s3client"

	"github.com/Diaku49/FoodOrderSystem/backend/internals/db"
	"github.com/Diaku49/FoodOrderSystem/backend/internals/email"
	"github.com/Diaku49/FoodOrderSystem/backend/internals/router"
	"github.com/Diaku49/FoodOrderSystem/backend/mq"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	port := os.Getenv("PORT")

	err = s3client.InitS3Client()
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

	log.Println("Server running on port:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
