package repository

import (
	"context"
	"hrms/internal/onboarding/model"
)

func (r *OnboardingRepository) AddExperience(req model.ExperienceRequest) error {  

	query := `
	INSERT INTO employee_experience
	(employee_id,company_name,designation,employment_type,start_date,end_date,is_current,industry,description)
	VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)
	`

	_, err := r.DB.Exec(context.Background(), query,
		req.EmployeeID, req.CompanyName, req.Designation, req.EmploymentType,
		req.StartDate, req.EndDate, req.IsCurrent, req.Industry, req.Description,
	)

	return err
}

func (r *OnboardingRepository) GetExperience(employeeID int) ([]model.Experience, error) {

	query := `
	SELECT id,employee_id,company_name,designation,employment_type,
	       start_date,end_date,is_current,industry,description
	FROM employee_experience
	WHERE employee_id=$1
	`

	rows, err := r.DB.Query(context.Background(), query, employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.Experience

	for rows.Next() {
		var exp model.Experience

		err := rows.Scan(
			&exp.ID, &exp.EmployeeID, &exp.CompanyName, &exp.Designation,
			&exp.EmploymentType, &exp.StartDate, &exp.EndDate,
			&exp.IsCurrent, &exp.Industry, &exp.Description,
		)
		if err != nil {
			return nil, err
		}

		list = append(list, exp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

func (r *OnboardingRepository) UpdateExperience(id int, req model.ExperienceRequest) error {  

	query := `
	UPDATE employee_experience
	SET company_name=$1, designation=$2, employment_type=$3,
	    start_date=$4, end_date=$5, is_current=$6,
	    industry=$7, description=$8
	WHERE id=$9
	`

	_, err := r.DB.Exec(context.Background(), query,
		req.CompanyName, req.Designation, req.EmploymentType,
		req.StartDate, req.EndDate, req.IsCurrent,
		req.Industry, req.Description, id,
	)

	return err
}

func (r *OnboardingRepository) DeleteExperience(id int) error {
	_, err := r.DB.Exec(context.Background(),
		`DELETE FROM employee_experience WHERE id=$1`, id)
	return err
}