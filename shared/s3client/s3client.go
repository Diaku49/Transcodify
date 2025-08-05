package s3client

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

type StorageClient struct {
	Client    *s3.S3
	endpoint  string
	accessKey string
	secretKey string
}

var S3Client *StorageClient
var Bucket string

func InitS3Client() error {
	if S3Client == nil {
		// Try multiple paths for .env file
		envPaths := []string{
			".env",
			"../shared/s3client/.env",
			"../../shared/s3client/.env",
		}

		var envLoaded bool
		for _, path := range envPaths {
			if err := godotenv.Load(path); err == nil {
				log.Printf("Successfully loaded .env from: %s", path)
				envLoaded = true
				break
			}
		}

		if !envLoaded {
			log.Fatal("Error loading .env file from any location")
			return fmt.Errorf("could not load .env file")
		}

		endpoint := os.Getenv("LIARA_ENDPOINT_URL")
		accessKey := os.Getenv("LIARA_ACCESS_KEY")
		secretKey := os.Getenv("LIARA_SECRET_KEY")
		Bucket = os.Getenv("BUCKET_NAME")

		// Debug logging
		log.Printf("S3 Config - Endpoint: %s", endpoint)
		log.Printf("S3 Config - AccessKey: %s", accessKey)
		log.Printf("S3 Config - SecretKey: %s", secretKey)
		log.Printf("S3 Config - Bucket: %s", Bucket)

		sess, err := session.NewSession(&aws.Config{
			Endpoint:         aws.String(endpoint),
			Region:           aws.String("us-east-1"),
			Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
			S3ForcePathStyle: aws.Bool(true),
		})
		if err != nil {
			log.Fatalf("Failed to create AWS session: %v", err)
			return err
		}

		sc := s3.New(sess)

		S3Client = &StorageClient{
			Client:    sc,
			secretKey: secretKey,
			accessKey: accessKey,
			endpoint:  endpoint,
		}
		return nil
	}
	return nil
}
