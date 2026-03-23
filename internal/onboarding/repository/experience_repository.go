package repository

import "hrms/internal/onboarding/model"

func (r *OnboardingRepository) AddExperience(exp model.Experience) error {

	query := `
	INSERT INTO employee_experience
	(employee_id,company,role,start_date,end_date,currently_working)
	VALUES($1,$2,$3,$4,$5,$6)
	`

	_, err := r.DB.Exec(
		query,
		exp.EmployeeID,
		exp.Company,
		exp.Role,
		exp.StartDate,
		exp.EndDate,
		exp.CurrentlyWorking,
	)

	return err
}

func (r *OnboardingRepository) GetExperience(employeeID int) ([]model.Experience, error) {

	query := `
	SELECT id,employee_id,company,role,start_date,end_date,currently_working
	FROM employee_experience
	WHERE employee_id=$1
	`

	rows, err := r.DB.Query(query, employeeID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var list []model.Experience

	for rows.Next() {

		var exp model.Experience

		err := rows.Scan(
			&exp.ID,
			&exp.EmployeeID,
			&exp.Company,
			&exp.Role,
			&exp.StartDate,
			&exp.EndDate,
			&exp.CurrentlyWorking,
		)

		if err != nil {
			return nil, err
		}

		list = append(list, exp)
	}

	return list, nil
}

func (r *OnboardingRepository) UpdateExperience(id int, exp model.Experience) error {

	query := `
	UPDATE employee_experience
	SET company=$1,
	    role=$2,
	    start_date=$3,
	    end_date=$4,
	    currently_working=$5
	WHERE id=$6
	`

	_, err := r.DB.Exec(
		query,
		exp.Company,
		exp.Role,
		exp.StartDate,
		exp.EndDate,
		exp.CurrentlyWorking,
		id,
	)

	return err
}

func (r *OnboardingRepository) DeleteExperience(id int) error {

	query := `
	DELETE FROM employee_experience
	WHERE id=$1
	`

	_, err := r.DB.Exec(query, id)

	return err
}