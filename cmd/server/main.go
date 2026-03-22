package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"hrms/config"
	"hrms/internal/onboarding/routes"
	"hrms/pkg/database"
	"hrms/pkg/middleware"
)

func main() {

	// load env
	config.LoadEnv()

	// connect db
	err := database.ConnectDB()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	err = database.RunMigrations(database.DB)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	// init logger
	middleware.InitLogger()

	r := gin.Default()

	// middleware
	r.Use(middleware.ErrorHandler())

	api := r.Group("/api")

	// register module routes
	routes.RegisterOnboardingRoutes(api)

	port := os.Getenv("PORT")

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
