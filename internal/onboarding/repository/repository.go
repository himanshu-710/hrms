package repository

import "github.com/jackc/pgx/v5/pgxpool"

type OnboardingRepository struct {
	DB *pgxpool.Pool
}

func NewOnboardingRepository(db *pgxpool.Pool) *OnboardingRepository {
	return &OnboardingRepository{
		DB: db,
	}
}