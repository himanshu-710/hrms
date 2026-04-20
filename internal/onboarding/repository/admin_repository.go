package repository

import (
	"context"
	"fmt"
	"hrms/internal/onboarding/model"
	"time"
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

func (r *OnboardingRepository) CreateEmployee(req model.CreateEmployeeRequest) (*model.Employee, error) {
	// Generate employee code (you can customize this logic)
	employeeCode := fmt.Sprintf("EMP%06d", time.Now().Unix()%1000000)

	query := `
	INSERT INTO employees
	(employee_code, work_email, first_name, last_name, personal_email, department, employment_context_role, date_of_joining)
	VALUES($1, $2, $3, $4, $5, $6, $7, CURRENT_DATE)
	RETURNING id, employee_code, work_email, first_name, last_name, employment_context_role
	`

	var emp model.Employee
	err := r.DB.QueryRow(context.Background(), query,
		employeeCode, req.Email, req.FirstName, req.LastName,
		req.Email, req.Department, req.EmploymentContextRole,
	).Scan(&emp.ID, &emp.EmployeeCode, &emp.WorkEmail, &emp.FirstName, &emp.LastName, &emp.EmploymentContextRole)

	if err != nil {
		return nil, err
	}

	return &emp, nil
}
