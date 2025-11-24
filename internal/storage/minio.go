package storage

import (
	// "context"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client
var BucketName string
var ServerURL string

func InitMinio() {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	BucketName = os.Getenv("MINIO_BUCKET")
	useSSL := os.Getenv("MINIO_USE_SSL") == "true"
	ServerURL = os.Getenv("SERVER_URL")
	if ServerURL == "" {
		ServerURL = "http://localhost:8080"
	}

	var err error
	MinioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln("Failed to connect to MinIO:", err)
	}

	log.Println("MinIO connected successfully")
}
