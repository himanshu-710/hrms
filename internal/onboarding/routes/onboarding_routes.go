package routes

import (
	"github.com/gofiber/fiber/v2"

	"hrms/internal/onboarding/handler"
	"hrms/internal/onboarding/repository"
	"hrms/internal/onboarding/service"
	"hrms/pkg/database"
	"hrms/pkg/storage" 
)

func RegisterOnboardingRoutes(app *fiber.App) {

	repo := repository.NewOnboardingRepository(database.DB)
	store :=  storage.NewLocalStorage()  
	service := service.NewOnboardingService(repo,store)
	handler := handler.NewOnboardingHandler(service)

	onboarding := app.Group("/api/v1/onboarding")

	// Health check
	onboarding.Get("/health", handler.Health)

	// Employee Profile
	onboarding.Post("/employee", handler.CreateEmployee)
	onboarding.Get("/profile/:id", handler.GetProfile)

	onboarding.Put("/profile/primary", handler.UpdatePrimaryDetails)
	onboarding.Put("/profile/contact", handler.UpdateContactDetails)
	onboarding.Put("/relations", handler.UpdateRelations)

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

	//completion
	onboarding.Get("/completion", handler.GetCompletion)

	//idnetity

	onboarding.Put("/identity", handler.SaveIdentity)
	onboarding.Get("/identity", handler.GetIdentity)

	//docs
	onboarding.Post("/documents/upload", handler.UploadDocument)
onboarding.Get("/documents", handler.GetDocuments)
onboarding.Delete("/documents/:id", handler.DeleteDocument)
onboarding.Patch("/documents/:id/verify", handler.VerifyDocument)

//admin
onboarding.Get("/admin/dashboard", handler.GetDashboard)
//assets
onboarding.Get("/assets", handler.GetAssets)
onboarding.Patch("/assets/:id/acknowledge", handler.AcknowledgeAsset)
onboarding.Post("/assets", handler.AssignAsset)

}
