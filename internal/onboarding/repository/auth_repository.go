package repository

import (
	"context"
	"hrms/internal/onboarding/model"
)

func (r *OnboardingRepository) GetEmployeeByCode(code string) (*model.AuthEmployee, error) {

	query := `
	SELECT id, work_email, employee_code, 
	       COALESCE(password_hash, ''), 
	       COALESCE(employment_context_role, 'EMPLOYEE')
	FROM employees
	WHERE employee_code = $1 AND is_active = true
	`

	var emp model.AuthEmployee
	err := r.DB.QueryRow(context.Background(), query, code).Scan(
		&emp.ID,
		&emp.WorkEmail,
		&emp.EmployeeCode,
		&emp.PasswordHash,
		&emp.Role,
	)
	if err != nil {
		return nil, err
	}

	return &emp, nil
}

func (r *OnboardingRepository) GetEmployeeByWorkEmail(email string) (*model.AuthEmployee, error) {

	query := `
	SELECT id, work_email, employee_code, 
	       COALESCE(password_hash, ''), 
	       COALESCE(employment_context_role, 'EMPLOYEE')
	FROM employees
	WHERE work_email = $1 AND is_active = true
	`

	var emp model.AuthEmployee
	err := r.DB.QueryRow(context.Background(), query, email).Scan(
		&emp.ID,
		&emp.WorkEmail,
		&emp.EmployeeCode,
		&emp.PasswordHash,
		&emp.Role,
	)
	if err != nil {
		return nil, err
	}

	return &emp, nil
}

func (r *OnboardingRepository) SetPasswordHash(employeeID int, hash string) error {
	_, err := r.DB.Exec(context.Background(),
		`UPDATE employees SET password_hash = $1 WHERE id = $2`,
		hash, employeeID,
	)
	return err
}

func (r *OnboardingRepository) StoreRefreshToken(employeeID int, tokenHash string, expiry string) error {
	_, err := r.DB.Exec(context.Background(),
		`UPDATE employees 
		 SET refresh_token_hash = $1, refresh_token_expiry = $2 
		 WHERE id = $3`,
		tokenHash, expiry, employeeID,
	)
	return err
}

func (r *OnboardingRepository) GetRefreshTokenData(employeeID int) (string, string, error) {
	var hash, expiry string
	err := r.DB.QueryRow(context.Background(),
		`SELECT refresh_token_hash, refresh_token_expiry::TEXT 
		 FROM employees WHERE id = $1`,
		employeeID,
	).Scan(&hash, &expiry)
	return hash, expiry, err
}

func (r *OnboardingRepository) ClearRefreshToken(employeeID int) error {
	_, err := r.DB.Exec(context.Background(),
		`UPDATE employees 
		 SET refresh_token_hash = NULL, refresh_token_expiry = NULL 
		 WHERE id = $1`,
		employeeID,
	)
	return err
}
func (r *OnboardingRepository) GetEmployeeByRefreshHash(hash string) (*model.AuthEmployee, string, error) {

	query := `
	SELECT id, work_email, employee_code, 
	       COALESCE(password_hash, ''),
	       COALESCE(employment_context_role, 'EMPLOYEE'),
	       refresh_token_expiry::TEXT
	FROM employees
	WHERE refresh_token_hash = $1 AND is_active = true
	`

	var emp model.AuthEmployee
	var expiry string

	err := r.DB.QueryRow(context.Background(), query, hash).Scan(
		&emp.ID,
		&emp.WorkEmail,
		&emp.EmployeeCode,
		&emp.PasswordHash,
		&emp.Role,
		&expiry,
	)
	if err != nil {
		return nil, "", err
	}

	return &emp, expiry, nil
}
