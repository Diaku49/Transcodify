package s3client

import (
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
	Bucket    string
}

var S3Client *StorageClient
var Bucket string

func InitS3Client() error {
	if S3Client == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
			return err
		}

		endpoint := os.Getenv("LIARA_ENDPOINT_URL")
		accessKey := os.Getenv("LIARA_ACCESS_KEY")
		secretKey := os.Getenv("LIARA_SECRET_KEY")
		Bucket = os.Getenv("BUCKET_NAME")

		sess, err := session.NewSession(&aws.Config{
			Endpoint:    aws.String(endpoint),
			Region:      aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
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
			Bucket:    Bucket,
		}
		return nil
	}
	return nil
}
