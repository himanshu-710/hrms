package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	"hrms/config"
	"hrms/internal/onboarding/routes"
	"hrms/pkg/database"
	"hrms/pkg/middleware"
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

	routes.RegisterOnboardingRoutes(app)

	port := os.Getenv("PORT")

	log.Fatal(app.Listen(":" + port))
}