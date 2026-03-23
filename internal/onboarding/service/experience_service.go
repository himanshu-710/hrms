package service

import "hrms/internal/onboarding/model"

func (s *OnboardingService) AddExperience(exp model.Experience) error {
	return s.Repo.AddExperience(exp)
}

func (s *OnboardingService) GetExperience(employeeID int) ([]model.Experience, error) {
	return s.Repo.GetExperience(employeeID)
}

func (s *OnboardingService) UpdateExperience(id int, exp model.Experience) error {
	return s.Repo.UpdateExperience(id, exp)
}

func (s *OnboardingService) DeleteExperience(id int) error {
	return s.Repo.DeleteExperience(id)
}