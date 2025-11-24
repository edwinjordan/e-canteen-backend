package config

import (
	"log"
	"os"
	"strconv"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioClient() *minio.Client {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	useSSLStr := os.Getenv("MINIO_USE_SSL")

	useSSL, err := strconv.ParseBool(useSSLStr)
	if err != nil {
		useSSL = false // Default to false if not specified or invalid
	}

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully connected to MinIO at %s", endpoint)

	// Optional: Check if bucket exists, create if not?
	// For now, just return the client.

	return minioClient
}
