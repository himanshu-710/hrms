package service

import "hrms/internal/onboarding/model"

func (s *OnboardingService) AddEducation(req model.EducationRequest) error { 
	return s.Repo.AddEducation(req)
}

func (s *OnboardingService) GetEducation(employeeID int) ([]model.Education, error) {
	return s.Repo.GetEducation(employeeID)
}

func (s *OnboardingService) UpdateEducation(id int, req model.EducationRequest) error {  
	return s.Repo.UpdateEducation(id, req)
}

func (s *OnboardingService) DeleteEducation(id int) error {
	return s.Repo.DeleteEducation(id)
}