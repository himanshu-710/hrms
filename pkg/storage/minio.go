package storage

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinIOStorage struct {
	client *minio.Client
	bucket string
}

func NewMinIOStorage() (*MinIOStorage, error) {

	client, err := minio.New("localhost:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("minioadmin", "minioadmin", ""),
		Secure: false,
	})

	if err != nil {
		return nil, err
	}

	return &MinIOStorage{
		client: client,
		bucket: "hrms-docs",
	}, nil
}

func (m *MinIOStorage) Upload(file *multipart.FileHeader, path string) (string, error) {

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	_, err = m.client.PutObject(
		context.Background(),
		m.bucket,
		path,
		src,
		file.Size,
		minio.PutObjectOptions{
			ContentType: file.Header.Get("Content-Type"),
		},
	)

	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("http://localhost:9000/%s/%s", m.bucket, path)

	return url, nil
}
func (m *MinIOStorage) Delete(path string) error {

	return m.client.RemoveObject(
		context.Background(),
		m.bucket,
		path,
		minio.RemoveObjectOptions{},
	)
}