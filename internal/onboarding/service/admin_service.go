package service

import (
	"time"

	"hrms/internal/onboarding/model"
)

func (s *OnboardingService) GetDashboard() ([]model.DashboardRowDTO, error) {

	employees, err := s.Repo.GetAllEmployees()
	if err != nil {
		return nil, err
	}

	var result []model.DashboardRowDTO

	for _, emp := range employees {

		comp, err := s.ComputeCompletion(emp.ID)
		if err != nil {
			return nil, err
		}

		// find incomplete sections
		var incomplete []string
		for k, v := range comp.Sections {
			if !v {
				incomplete = append(incomplete, k)
			}
		}

		// days since joining
		var days int
		if emp.DateOfJoining != nil {
			days = int(time.Since(*emp.DateOfJoining).Hours() / 24)
		}

		result = append(result, model.DashboardRowDTO{
			EmployeeID:         emp.ID,
			Name:               emp.FirstName + " " + emp.LastName,
			DaysSinceJoining:   days,
			CompletionPct:      comp.Percentage,
			IncompleteSections: incomplete,
		})
	}

	return result, nil
}