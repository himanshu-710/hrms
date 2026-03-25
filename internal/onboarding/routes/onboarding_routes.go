package routes

import (
	"github.com/gofiber/fiber/v2"

	"hrms/internal/onboarding/handler"
	"hrms/internal/onboarding/repository"
	"hrms/internal/onboarding/service"
	"hrms/pkg/database"
)

func RegisterOnboardingRoutes(app *fiber.App) {

	repo := repository.NewOnboardingRepository(database.DB)
	service := service.NewOnboardingService(repo)
	handler := handler.NewOnboardingHandler(service)

	onboarding := app.Group("/api/onboarding")

	// Health check
	onboarding.Get("/health", handler.Health)

	// Employee Profile
	onboarding.Post("/employee", handler.CreateEmployee)
	onboarding.Get("/profile/:id", handler.GetProfile)
	onboarding.Put("/profile/:id", handler.UpdateProfile)

	// Education
	onboarding.Post("/education", handler.AddEducation)
	onboarding.Get("/education/:employeeId", handler.GetEducation)
	onboarding.Put("/education/:id", handler.UpdateEducation)
	onboarding.Delete("/education/:id", handler.DeleteEducation)

	// Experience
	onboarding.Post("/experience", handler.AddExperience)
	onboarding.Get("/experience/:employeeId", handler.GetExperience)
	onboarding.Put("/experience/:id", handler.UpdateExperience)
	onboarding.Delete("/experience/:id", handler.DeleteExperience)

	// Addresses
	onboarding.Put("/addresses", handler.SaveAddresses)
}