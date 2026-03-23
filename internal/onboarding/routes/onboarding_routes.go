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

	api := app.Group("/api")
	onboarding := api.Group("/onboarding")

	onboarding.Get("/health", handler.Health)

	onboarding.Post("/employee", handler.CreateEmployee)

	onboarding.Post("/profile", handler.CreateEmployee)
	onboarding.Get("/profile/:id", handler.GetProfile)

	onboarding.Post("/education", handler.AddEducation)
	onboarding.Get("/education/:employeeId", handler.GetEducation)
	onboarding.Delete("/education/:id", handler.DeleteEducation)

	onboarding.Post("/experience", handler.AddExperience)
	onboarding.Get("/experience/:employeeId", handler.GetExperience)
	onboarding.Delete("/experience/:id", handler.DeleteExperience)

	onboarding.Put("/addresses", handler.SaveAddresses)
}