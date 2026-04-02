package repository

import (
	"context"
	"hrms/internal/onboarding/model"
)

func (r *OnboardingRepository) AddEducation(req model.EducationRequest) error {  

	query := `
	INSERT INTO employee_education
	(employee_id,degree,branch,university,cgpa_or_pct,year_of_joining,year_of_completion)
	VALUES($1,$2,$3,$4,$5,$6,$7)
	`

	_, err := r.DB.Exec(context.Background(), query,
		req.EmployeeID, req.Degree, req.Branch, req.University,
		req.CGPAOrPct, req.YearOfJoining, req.YearOfCompletion,
	)

	return err
}

func (r *OnboardingRepository) GetEducation(employeeID int) ([]model.Education, error) {

	query := `
	SELECT id,employee_id,degree,branch,university,cgpa_or_pct,year_of_joining,year_of_completion
	FROM employee_education
	WHERE employee_id=$1
	`

	rows, err := r.DB.Query(context.Background(), query, employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.Education

	for rows.Next() {
		var edu model.Education

		err := rows.Scan(
			&edu.ID, &edu.EmployeeID, &edu.Degree, &edu.Branch,
			&edu.University, &edu.CGPAOrPct,
			&edu.YearOfJoining, &edu.YearOfCompletion,
		)
		if err != nil {
			return nil, err
		}

		list = append(list, edu)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

func (r *OnboardingRepository) UpdateEducation(id int, req model.EducationRequest) error {  // changed

	query := `
	UPDATE employee_education
	SET degree=$1, branch=$2, university=$3, cgpa_or_pct=$4,
	    year_of_joining=$5, year_of_completion=$6
	WHERE id=$7
	`

	_, err := r.DB.Exec(context.Background(), query,
		req.Degree, req.Branch, req.University, req.CGPAOrPct,
		req.YearOfJoining, req.YearOfCompletion, id,
	)

	return err
}

func (r *OnboardingRepository) DeleteEducation(id int) error {
	_, err := r.DB.Exec(context.Background(),
		`DELETE FROM employee_education WHERE id=$1`, id)
	return err
}
func (r *OnboardingRepository) GetEducationOwner(id int) (int, error) {
	var employeeID int
	err := r.DB.QueryRow(context.Background(),
		`SELECT employee_id FROM employee_education WHERE id=$1`, id,
	).Scan(&employeeID)
	return employeeID, err
}