package storage

import (
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type LocalStorage struct{}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{}
}

func (l *LocalStorage) Upload(file *multipart.FileHeader, path string) (string, error) {
	if file == nil {
		return "", fmt.Errorf("file is required")
	}
	if file.Size > 5*1024*1024 {
		return "", fmt.Errorf("file size exceeds 5MB limit")
	}

	allowedTypes := map[string]bool{
		"application/pdf": true,
		"image/jpeg":      true,
		"image/png":       true,
	}
	ext := filepath.Ext(file.Filename)
	mimeType := mime.TypeByExtension(ext)
	if !allowedTypes[mimeType] {
		return "", fmt.Errorf("unsupported file type")
	}
	if strings.Contains(path, "..") {
		return "", fmt.Errorf("invalid file path")
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return path, nil
}

func (l *LocalStorage) Delete(path string) error {
	if path == "" {
		return fmt.Errorf("invalid file path")
	}
	return os.Remove(path)
}

func (l *LocalStorage) GetPresignedURL(objectPath string, expiry time.Duration) (string, error) {
	_ = expiry

	if objectPath == "" {
		return "", fmt.Errorf("object path is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return fmt.Sprintf("http://localhost:%s/%s", port, strings.TrimLeft(filepath.ToSlash(objectPath), "/")), nil
}
