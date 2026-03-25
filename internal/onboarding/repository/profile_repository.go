package repository

import "hrms/internal/onboarding/model"

func (r *OnboardingRepository) CreateEmployee(firstName string, lastName string, email string, department string) error {

	query := `
	INSERT INTO employees(first_name,last_name,personal_email,department)
	VALUES($1,$2,$3,$4)
	`

	_, err := r.DB.Exec(query, firstName, lastName, email, department)

	return err
}

func (r *OnboardingRepository) GetEmployee(id int) (*model.Employee, error) {

	query := `
	SELECT id,first_name,last_name,personal_email,department
	FROM employees
	WHERE id=$1
	`

	row := r.DB.QueryRow(query, id)

	var emp model.Employee

	err := row.Scan(
		&emp.ID,
		&emp.FirstName,
		&emp.LastName,
		&emp.PersonalEmail,
		&emp.Department,
	)

	if err != nil {
		return nil, err
	}

	return &emp, nil
}

func (r *OnboardingRepository) UpdateEmployee(id int, firstName string, lastName string, email string, department string) error {

	query := `
	UPDATE employees
	SET first_name=$1,
	    last_name=$2,
	    personal_email=$3,
	    department=$4
	WHERE id=$5
	`

	_, err := r.DB.Exec(query, firstName, lastName, email, department, id)

	return err
}