package service

import (
	"hrms/internal/onboarding/model"
	"hrms/internal/onboarding/repository"
)

type OnboardingService struct {
	Repo *repository.OnboardingRepository
}

func NewOnboardingService(repo *repository.OnboardingRepository) *OnboardingService {
	return &OnboardingService{
		Repo: repo,
	}
}

func (s *OnboardingService) CreateEmployee(firstName string, lastName string, email string, department string) error {
	return s.Repo.CreateEmployee(firstName, lastName, email, department)
}

func (s *OnboardingService) GetEmployee(id int) (*model.Employee, error) {
	return s.Repo.GetEmployee(id)
}

func (s *OnboardingService) AddEducation(edu model.Education) error {
	return s.Repo.AddEducation(edu)
}

func (s *OnboardingService) GetEducation(employeeID int) ([]model.Education, error) {
	return s.Repo.GetEducation(employeeID)
}

func (s *OnboardingService) DeleteEducation(id int) error {
	return s.Repo.DeleteEducation(id)
}

func (s *OnboardingService) AddExperience(exp model.Experience) error {
	return s.Repo.AddExperience(exp)
}

func (s *OnboardingService) GetExperience(employeeID int) ([]model.Experience, error) {
	return s.Repo.GetExperience(employeeID)
}

func (s *OnboardingService) DeleteExperience(id int) error {
	return s.Repo.DeleteExperience(id)
}
func (s *OnboardingService) SaveAddresses(req model.AddressesRequest) error {

	if req.CopyFromCurrent {
		req.Permanent = req.Current
	}

	return s.Repo.SaveAddresses(req)
}
