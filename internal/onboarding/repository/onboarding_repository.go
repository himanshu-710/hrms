package repository

import "database/sql"

type OnboardingRepository struct {
	DB *sql.DB
}

func NewOnboardingRepository(db *sql.DB) *OnboardingRepository {
	return &OnboardingRepository{
		DB: db,
	}
}

func (r *OnboardingRepository) CreateEmployee(name string, email string) error {

	query := `
	INSERT INTO employees(name,email)
	VALUES($1,$2)
	`

	_, err := r.DB.Exec(query, name, email)

	return err
}