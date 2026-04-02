package storage

import (
    "mime/multipart"
    "time"
)

type Storage interface {
    Upload(file *multipart.FileHeader, path string) (string, error)
    Delete(path string) error
    GetPresignedURL(objectPath string, expiry time.Duration) (string, error)
}