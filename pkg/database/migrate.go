package database

import (
	"context"
	"fmt"
	"os"
)

func RunMigrations() error {

	file, err := os.ReadFile("migrations/003_create_onboarding_tables.sql")
	if err != nil {
		return err
	}

	_, err = DB.Exec(context.Background(), string(file))
	if err != nil {
		return err
	}

	fmt.Println("Database migrations applied")
	return nil
}