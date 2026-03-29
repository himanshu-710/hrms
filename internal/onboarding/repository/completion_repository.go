package repository

import "context"

func (r *OnboardingRepository) HasEducation(employeeID int) (bool, error) {
	var count int
	err := r.DB.QueryRow(context.Background(),
		`SELECT COUNT(*) FROM employee_education WHERE employee_id=$1`, employeeID).Scan(&count)
	return count > 0, err
}

func (r *OnboardingRepository) HasExperience(employeeID int) (bool, error) {
	var count int
	err := r.DB.QueryRow(context.Background(),
		`SELECT COUNT(*) FROM employee_experience WHERE employee_id=$1`, employeeID).Scan(&count)
	return count > 0, err
}

func (r *OnboardingRepository) HasAddress(employeeID int) (bool, error) {
	var count int
	err := r.DB.QueryRow(context.Background(),
		`SELECT COUNT(*) FROM employee_addresses WHERE employee_id=$1`, employeeID).Scan(&count)
	return count > 0, err
}

func (r *OnboardingRepository) HasDocuments(employeeID int) (bool, error) {
	var count int
	err := r.DB.QueryRow(context.Background(),
		`SELECT COUNT(*) FROM employee_documents WHERE employee_id=$1`, employeeID).Scan(&count)
	return count > 0, err
}

func (r *OnboardingRepository) HasIdentity(employeeID int) (bool, error) {
	var count int
	err := r.DB.QueryRow(context.Background(),
		`SELECT COUNT(*) FROM employee_identity_documents WHERE employee_id=$1`, employeeID).Scan(&count)
	return count > 0, err
}

func (r *OnboardingRepository) HasAssets(employeeID int) (bool, error) {
	var count int
	err := r.DB.QueryRow(context.Background(),
		`SELECT COUNT(*) FROM employee_assets WHERE employee_id=$1 AND is_active=true`, employeeID).Scan(&count)
	return count > 0, err
}