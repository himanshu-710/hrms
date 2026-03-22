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

func (s *OnboardingService) CreateEmployee(name string, email string) error {
	return s.Repo.CreateEmployee(name, email)
}