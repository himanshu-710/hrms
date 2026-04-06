package service

import "hrms/internal/onboarding/model"

func (s *OnboardingService) CreateEmployee(firstName string, lastName string, email string, department string, employmentContextRole string) error {
	if employmentContextRole == "" {
		employmentContextRole = "EMPLOYEE"
	}

	return s.Repo.CreateEmployee(firstName, lastName, email, department, employmentContextRole)
}

func (s *OnboardingService) GetEmployee(id int) (*model.Employee, error) {
	return s.Repo.GetEmployee(id)
}

func (s *OnboardingService) GetFullProfile(id int) (*model.OnboardingProfileDTO, error) {
	return s.Repo.GetFullProfile(id)
}

func (s *OnboardingService) UpdatePrimaryDetails(id int, req model.PrimaryDetailsRequest) error {
	return s.Repo.UpdatePrimaryDetails(id, req)
}

func (s *OnboardingService) UpdateContactDetails(id int, req model.ContactRequest) error {
	return s.Repo.UpdateContactDetails(id, req)
}

func (s *OnboardingService) UpdateRelations(id int, relations map[string]interface{}) error {
	return s.Repo.UpdateRelations(id, relations)
}
