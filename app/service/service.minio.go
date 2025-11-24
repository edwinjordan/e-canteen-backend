package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type MinioService interface {
	UploadFile(ctx context.Context, fileHeader *multipart.FileHeader, folder string) (string, error)
	GetFileUrl(ctx context.Context, objectName string) (string, error)
}

type MinioServiceImpl struct {
	Client     *minio.Client
	BucketName string
}

func NewMinioService(client *minio.Client) MinioService {
	bucketName := os.Getenv("MINIO_BUCKET_NAME")
	if bucketName == "" {
		bucketName = "e-canteen" // Default bucket name
	}

	// Ensure bucket exists
	ctx := context.Background()
	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		fmt.Printf("Error checking bucket existence: %v\n", err)
	} else if !exists {
		err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			fmt.Printf("Error creating bucket: %v\n", err)
		} else {
			fmt.Printf("Successfully created bucket %s\n", bucketName)
		}
	}

	// Set bucket policy to public read-only
	policy := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Principal": {"AWS": ["*"]},
				"Action": ["s3:GetObject"],
				"Resource": ["arn:aws:s3:::%s/*"]
			}
		]
	}`, bucketName)

	err = client.SetBucketPolicy(ctx, bucketName, policy)
	if err != nil {
		fmt.Printf("Error setting bucket policy: %v\n", err)
	}

	return &MinioServiceImpl{
		Client:     client,
		BucketName: bucketName,
	}
}

func (s *MinioServiceImpl) UploadFile(ctx context.Context, fileHeader *multipart.FileHeader, folder string) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Generate a unique file name
	ext := filepath.Ext(fileHeader.Filename)
	filename := fmt.Sprintf("%s/%s%s", folder, uuid.New().String(), ext)

	// Upload the file
	info, err := s.Client.PutObject(ctx, s.BucketName, filename, file, fileHeader.Size, minio.PutObjectOptions{
		ContentType: fileHeader.Header.Get("Content-Type"),
	})
	if err != nil {
		return "", err
	}

	fmt.Printf("Successfully uploaded %s of size %d\n", filename, info.Size)

	return filename, nil
}

func (s *MinioServiceImpl) GetFileUrl(ctx context.Context, objectName string) (string, error) {
	// For now, we can return the presigned URL or just the public URL if the bucket is public.
	// Let's return a presigned URL for safety, or just the path if we serve it via a proxy.
	// Assuming we want a direct link:

	// Check if we want a presigned URL
	// reqParams := make(url.Values)
	// presignedURL, err := s.Client.PresignedGetObject(ctx, s.BucketName, objectName, time.Hour*24, reqParams)
	// if err != nil {
	// 	return "", err
	// }
	// return presignedURL.String(), nil

	// If using MinIO directly exposed or via Nginx:
	endpoint := os.Getenv("MINIO_ENDPOINT")
	// If endpoint is internal (e.g. minio:9000), we might need an external URL env var.
	// For local dev, localhost:9000 is fine if exposed.

	// Simple construction for now:
	return fmt.Sprintf("http://%s/%s/%s", endpoint, s.BucketName, objectName), nil
}
