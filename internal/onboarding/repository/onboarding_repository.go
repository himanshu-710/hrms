package repository

import (
	"database/sql"
	"hrms/internal/onboarding/model"
)

type OnboardingRepository struct {
	DB *sql.DB
}

func NewOnboardingRepository(db *sql.DB) *OnboardingRepository {
	return &OnboardingRepository{
		DB: db,
	}
}

func (r *OnboardingRepository) CreateEmployee(firstName string, lastName string, email string, department string) error {

	query := 
	`INSERT INTO employees(first_name,last_name,personal_email,department)
	 VALUES($1,$2,$3,$4)`

	_, err := r.DB.Exec(query, firstName, lastName, email, department)

	return err
}

func (r *OnboardingRepository) GetEmployee(id int) (*model.Employee, error) {

	query := 
	`SELECT id,first_name,last_name,personal_email,department
	FROM employees
	WHERE id=$1`

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

func (r *OnboardingRepository) AddEducation(edu model.Education) error {

	query := `INSERT INTO employee_education
	(employee_id,degree,branch,university,cgpa_or_pct,year_of_joining,year_of_completion)
	VALUES($1,$2,$3,$4,$5,$6,$7)`

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

	query := `SELECT id,employee_id,degree,branch,university,cgpa_or_pct,year_of_joining,year_of_completion
	FROM employee_education
	WHERE employee_id=$1`

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

func (r *OnboardingRepository) DeleteEducation(id int) error {

	query := `DELETE FROM employee_education
	WHERE id=$1`

	_, err := r.DB.Exec(query, id)

	return err
}

func (r *OnboardingRepository) AddExperience(exp model.Experience) error {

	query := `INSERT INTO employee_experience
	(employee_id,company,role,start_date,end_date,currently_working)
	VALUES($1,$2,$3,$4,$5,$6)`

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

	query := `SELECT id,employee_id,company,role,start_date,end_date,currently_working
	FROM employee_experience
	WHERE employee_id=$1`

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

func (r *OnboardingRepository) DeleteExperience(id int) error {

	query := `
	DELETE FROM employee_experience
	WHERE id=$1
	`

	_, err := r.DB.Exec(query, id)

	return err
}
func (r *OnboardingRepository) SaveAddresses(req model.AddressesRequest) error {

	query := `
	INSERT INTO employee_addresses
	(employee_id,address_type,street,city,state,country,postal_code)
	VALUES($1,$2,$3,$4,$5,$6,$7)
	ON CONFLICT(employee_id,address_type)
	DO UPDATE SET
	street=EXCLUDED.street,
	city=EXCLUDED.city,
	state=EXCLUDED.state,
	country=EXCLUDED.country,
	postal_code=EXCLUDED.postal_code
	`

	_, err := r.DB.Exec(
		query,
		req.EmployeeID,
		"CURRENT",
		req.Current.Street,
		req.Current.City,
		req.Current.State,
		req.Current.Country,
		req.Current.PostalCode,
	)

	if err != nil {
		return err
	}

	_, err = r.DB.Exec(
		query,
		req.EmployeeID,
		"PERMANENT",
		req.Permanent.Street,
		req.Permanent.City,
		req.Permanent.State,
		req.Permanent.Country,
		req.Permanent.PostalCode,
	)

	return err
}
