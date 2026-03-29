package repository

import (
	"context"
	"hrms/internal/onboarding/model"
)

func (r *OnboardingRepository) GetAllEmployees() ([]model.Employee, error) {

	query := `
	SELECT id, first_name, last_name, date_of_joining
	FROM employees
	WHERE is_active = true
	`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.Employee

	for rows.Next() {
		var emp model.Employee

		err := rows.Scan(
			&emp.ID,
			&emp.FirstName,
			&emp.LastName,
			&emp.DateOfJoining, 
		)
		if err != nil {
			return nil, err
		}

		list = append(list, emp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}