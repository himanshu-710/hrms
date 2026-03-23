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

	
	config.LoadEnv()

	
	err := database.ConnectDB()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	err = database.RunMigrations(database.DB)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	middleware.InitLogger()

	r := gin.Default()


	r.Use(middleware.ErrorHandler())

	api := r.Group("/api")

	routes.RegisterOnboardingRoutes(api)

	port := os.Getenv("PORT")

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
