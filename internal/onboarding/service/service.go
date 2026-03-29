package service

import (
    "hrms/internal/onboarding/repository"
    "mime/multipart"
)

type StorageProvider interface {
    Upload(file *multipart.FileHeader, path string) (string, error)
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