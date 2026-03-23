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