package repository

import "hrms/internal/onboarding/model"

func (r *OnboardingRepository) AddEducation(edu model.Education) error {

	query := `
	INSERT INTO employee_education
	(employee_id,degree,branch,university,cgpa_or_pct,year_of_joining,year_of_completion)
	VALUES($1,$2,$3,$4,$5,$6,$7)
	`

	_, err := r.DB.Exec(
		query,
		edu.EmployeeID,
		edu.Degree,
		edu.Branch,
		edu.University,
		edu.CGPAOrPct,
		edu.YearOfJoining,
		edu.YearOfCompletion,
	)

	return err
}

func (r *OnboardingRepository) GetEducation(employeeID int) ([]model.Education, error) {

	query := `
	SELECT id,employee_id,degree,branch,university,cgpa_or_pct,year_of_joining,year_of_completion
	FROM employee_education
	WHERE employee_id=$1
	`

	rows, err := r.DB.Query(query, employeeID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var list []model.Education

	for rows.Next() {

		var edu model.Education

		err := rows.Scan(
			&edu.ID,
			&edu.EmployeeID,
			&edu.Degree,
			&edu.Branch,
			&edu.University,
			&edu.CGPAOrPct,
			&edu.YearOfJoining,
			&edu.YearOfCompletion,
		)

		if err != nil {
			return nil, err
		}

		list = append(list, edu)
	}

	return list, nil
}

func (r *OnboardingRepository) UpdateEducation(id int, edu model.Education) error {

	query := `
	UPDATE employee_education
	SET degree=$1,
	    branch=$2,
	    university=$3,
	    cgpa_or_pct=$4,
	    year_of_joining=$5,
	    year_of_completion=$6
	WHERE id=$7
	`

	_, err := r.DB.Exec(
		query,
		edu.Degree,
		edu.Branch,
		edu.University,
		edu.CGPAOrPct,
		edu.YearOfJoining,
		edu.YearOfCompletion,
		id,
	)

	return err
}

func (r *OnboardingRepository) DeleteEducation(id int) error {

	query := `
	DELETE FROM employee_education
	WHERE id=$1
	`

	_, err := r.DB.Exec(query, id)

	return err
}