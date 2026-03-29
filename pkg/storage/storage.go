package storage

import "mime/multipart"

type Storage interface {
	Upload(file *multipart.FileHeader, path string) (string, error)
	Delete(path string) error
}