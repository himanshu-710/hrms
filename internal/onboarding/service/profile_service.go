package service

import "hrms/internal/onboarding/model"

func (s *OnboardingService) CreateEmployee(firstName string, lastName string, email string, department string) error {
	return s.Repo.CreateEmployee(firstName, lastName, email, department)
}

func (s *OnboardingService) GetEmployee(id int) (*model.Employee, error) {
	return s.Repo.GetEmployee(id)
}