package service

import "hrms/internal/onboarding/repository"

type OnboardingService struct {
	Repo *repository.OnboardingRepository
}

func NewOnboardingService(repo *repository.OnboardingRepository) *OnboardingService {
	return &OnboardingService{
		Repo: repo,
	}
}