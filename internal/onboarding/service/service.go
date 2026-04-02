package service

import (
	"hrms/internal/onboarding/repository"

	"mime/multipart"
	"time"
)

type StorageProvider interface {
	Upload(file *multipart.FileHeader, path string) (string, error)
	Delete(path string) error
	GetPresignedURL(objectPath string, expiry time.Duration) (string, error)
}

type OnboardingService struct {
	Repo    *repository.OnboardingRepository
	Storage StorageProvider
}

func NewOnboardingService(repo *repository.OnboardingRepository, storage StorageProvider) *OnboardingService {
	return &OnboardingService{
		Repo:    repo,
		Storage: storage,
	}
}
