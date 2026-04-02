package storage

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/url"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinIOStorage struct {
	client   *minio.Client
	bucket   string
	endpoint string
}

func NewMinIOStorage() (*MinIOStorage, error) {

	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")
	bucket := os.Getenv("MINIO_BUCKET")

	if endpoint == "" {
		endpoint = "localhost:9000"
	}
	if bucket == "" {
		bucket = "hrms-docs"
	}

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	
	exists, err := client.BucketExists(context.Background(), bucket)
	if err != nil {
		return nil, err
	}
	if !exists {
		if err := client.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{}); err != nil {
			return nil, fmt.Errorf("could not create bucket: %w", err)
		}
	}

	return &MinIOStorage{
		client:   client,
		bucket:   bucket,
		endpoint: endpoint,
	}, nil
}


func (m *MinIOStorage) Upload(file *multipart.FileHeader, objectPath string) (string, error) {

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	_, err = m.client.PutObject(
		context.Background(),
		m.bucket,
		objectPath,
		src,
		file.Size,
		minio.PutObjectOptions{
			ContentType: file.Header.Get("Content-Type"),
		},
	)
	if err != nil {
		return "", err
	}

	
	return objectPath, nil
}

func (m *MinIOStorage) GetPresignedURL(objectPath string, expiry time.Duration) (string, error) {

	
	reqParams := make(url.Values)

	presignedURL, err := m.client.PresignedGetObject(
		context.Background(),
		m.bucket,
		objectPath,
		expiry,
		reqParams,
	)
	if err != nil {
		return "", fmt.Errorf("could not generate presigned URL: %w", err)
	}

	return presignedURL.String(), nil
}

func (m *MinIOStorage) Delete(objectPath string) error {
	return m.client.RemoveObject(
		context.Background(),
		m.bucket,
		objectPath,
		minio.RemoveObjectOptions{},
	)
}