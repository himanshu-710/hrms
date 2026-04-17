package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"hrms/config"
	"hrms/internal/onboarding/routes"
	"hrms/pkg/database"
	"hrms/pkg/middleware"
	"hrms/pkg/scheduler"
	"log"
	"os"
)

func main() {

	config.LoadEnv()

	err := database.ConnectDB()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	err = database.RunMigrations()
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	middleware.InitLogger()

	app := fiber.New()

	app.Use(middleware.ErrorHandler())
	app.Static("/uploads", "./uploads")

	onboardingService := routes.RegisterOnboardingRoutes(app)
	scheduler.StartOnboardingReminderCron(context.Background(), onboardingService)

	port := os.Getenv("PORT")

	log.Fatal(app.Listen(":" + port))
}
