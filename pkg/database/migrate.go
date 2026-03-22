package database

import (
	"database/sql"
	"fmt"
	"os"
)

func RunMigrations(db *sql.DB) error {

	file, err := os.ReadFile("migrations/001_create_onboarding_tables.sql")
	if err != nil {
		return err
	}

	query := string(file)

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	fmt.Println("Database migrations applied")

	return nil
}