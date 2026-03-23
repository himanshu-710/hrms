package service

import "hrms/internal/onboarding/model"

func (s *OnboardingService) AddEducation(edu model.Education) error {
	return s.Repo.AddEducation(edu)
}

func (s *OnboardingService) GetEducation(employeeID int) ([]model.Education, error) {
	return s.Repo.GetEducation(employeeID)
}

func (s *OnboardingService) UpdateEducation(id int, edu model.Education) error {
	return s.Repo.UpdateEducation(id, edu)
}

func (s *OnboardingService) DeleteEducation(id int) error {
	return s.Repo.DeleteEducation(id)
}