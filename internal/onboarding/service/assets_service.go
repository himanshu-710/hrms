package service

import "hrms/internal/onboarding/model"

func (s *OnboardingService) GetAssets(employeeID int) ([]model.EmployeeAsset, error) {
	return s.Repo.GetAssets(employeeID)
}

func (s *OnboardingService) AcknowledgeAsset(id int) error {
	return s.Repo.AcknowledgeAsset(id)
}

func (s *OnboardingService) AssignAsset(a model.EmployeeAsset) error {
	return s.Repo.AssignAsset(a)
}