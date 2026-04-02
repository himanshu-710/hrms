package routes

import (
	"strconv"

	"hrms/internal/onboarding/handler"
	"hrms/internal/onboarding/repository"
	"hrms/internal/onboarding/service"
	"hrms/pkg/database"
	"hrms/pkg/middleware"
	"hrms/pkg/storage"

	"github.com/gofiber/fiber/v2"
)

func RegisterOnboardingRoutes(app *fiber.App) {

	repo := repository.NewOnboardingRepository(database.DB)
	store := storage.NewLocalStorage()
	svc := service.NewOnboardingService(repo, store)
	h := handler.NewOnboardingHandler(svc)

	RegisterAuthRoutes(app, h)
	app.Post("/api/v1/onboarding/employee", h.CreateEmployee)
	onboarding := app.Group("/api/v1/onboarding", middleware.AuthMiddleware())

	onboarding.Get("/health", h.Health)

	// Employee

	onboarding.Get("/profile/:id", h.GetProfile)

	// Profile — employeeId in URL, simple ownership check
	onboarding.Put("/profile/:employeeId/primary", middleware.OwnershipGuard(), h.UpdatePrimaryDetails)
	onboarding.Put("/profile/:employeeId/contact", middleware.OwnershipGuard(), h.UpdateContactDetails)
	onboarding.Put("/profile/:employeeId/relations", middleware.OwnershipGuard(), h.UpdateRelations)
	onboarding.Put("/profile/:employeeId/addresses", middleware.OwnershipGuard(), h.SaveAddresses)

	// Education — ownerFn for update/delete because URL has :id not :employeeId
	onboarding.Post("/education", h.AddEducation)
	onboarding.Get("/education/:employeeId", middleware.OwnershipGuard(), h.GetEducation)
	onboarding.Put("/education/:id", middleware.OwnershipGuard(func(c *fiber.Ctx) (int, error) {
		id, _ := strconv.Atoi(c.Params("id"))
		return repo.GetEducationOwner(id)
	}), h.UpdateEducation)
	onboarding.Delete("/education/:id", middleware.OwnershipGuard(func(c *fiber.Ctx) (int, error) {
		id, _ := strconv.Atoi(c.Params("id"))
		return repo.GetEducationOwner(id)
	}), h.DeleteEducation)

	// Experience — same pattern
	onboarding.Post("/experience", h.AddExperience)
	onboarding.Get("/experience/:employeeId", middleware.OwnershipGuard(), h.GetExperience)
	onboarding.Put("/experience/:id", middleware.OwnershipGuard(func(c *fiber.Ctx) (int, error) {
		id, _ := strconv.Atoi(c.Params("id"))
		return repo.GetExperienceOwner(id)
	}), h.UpdateExperience)
	onboarding.Delete("/experience/:id", middleware.OwnershipGuard(func(c *fiber.Ctx) (int, error) {
		id, _ := strconv.Atoi(c.Params("id"))
		return repo.GetExperienceOwner(id)
	}), h.DeleteExperience)

	// Identity
	onboarding.Put("/profile/:employeeId/identity", middleware.OwnershipGuard(), h.SaveIdentity)
	onboarding.Get("/profile/:employeeId/identity", middleware.OwnershipGuard(), h.GetIdentity)

	// Documents
	onboarding.Post("/profile/:employeeId/documents", middleware.OwnershipGuard(), h.UploadDocument)
	onboarding.Get("/profile/:employeeId/documents", middleware.OwnershipGuard(), h.GetDocuments)
	onboarding.Delete("/documents/:id", middleware.OwnershipGuard(func(c *fiber.Ctx) (int, error) {
		id, _ := strconv.Atoi(c.Params("id"))
		return repo.GetDocumentOwner(id)
	}), h.DeleteDocument)
	onboarding.Patch("/documents/:id/verify", h.VerifyDocument) // HR check in service layer

	// Assets
	onboarding.Post("/profile/:employeeId/assets", h.AssignAsset)
	onboarding.Get("/profile/:employeeId/assets", middleware.OwnershipGuard(), h.GetAssets)
	onboarding.Patch("/assets/:id/acknowledge", middleware.OwnershipGuard(func(c *fiber.Ctx) (int, error) {
		id, _ := strconv.Atoi(c.Params("id"))
		return repo.GetAssetOwner(id)
	}), h.AcknowledgeAsset)

	// Completion & Admin
	onboarding.Get("/profile/:employeeId/completion", middleware.OwnershipGuard(), h.GetCompletion)
	onboarding.Get("/admin/dashboard", h.GetDashboard)
}
