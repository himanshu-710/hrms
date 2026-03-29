package service

import "hrms/internal/onboarding/model"

func (s *OnboardingService) AddExperience(req model.ExperienceRequest) error {  // changed
	return s.Repo.AddExperience(req)
}

func (s *OnboardingService) GetExperience(employeeID int) ([]model.Experience, error) {
	return s.Repo.GetExperience(employeeID)
}

func (s *OnboardingService) UpdateExperience(id int, req model.ExperienceRequest) error {  // changed
	return s.Repo.UpdateExperience(id, req)
}

func (s *OnboardingService) DeleteExperience(id int) error {
	return s.Repo.DeleteExperience(id)
}