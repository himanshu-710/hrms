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
	store, err := storage.NewMinIOStorage()
	if err != nil {
		panic(err)
	}
	svc := service.NewOnboardingService(repo, store)
	h := handler.NewOnboardingHandler(svc)
	hrOnly := middleware.RequireRoles("HR")

	RegisterAuthRoutes(app, h)
	app.Get("/api/v1/onboarding/health", middleware.AuthMiddleware(), hrOnly, h.Health)
	app.Post("/api/v1/onboarding/employee", h.CreateEmployee)

	onboarding := app.Group("/api/v1/onboarding", middleware.AuthMiddleware())

	onboarding.Get("/profile/:id", middleware.OwnershipGuard(func(c *fiber.Ctx) (int, error) {
		return strconv.Atoi(c.Params("id"))
	}), h.GetProfile)

	onboarding.Put("/profile/:employeeId/primary", middleware.OwnershipGuard(), h.UpdatePrimaryDetails)
	onboarding.Put("/profile/:employeeId/contact", middleware.OwnershipGuard(), h.UpdateContactDetails)
	onboarding.Put("/profile/:employeeId/relations", middleware.OwnershipGuard(), h.UpdateRelations)
	onboarding.Put("/profile/:employeeId/addresses", middleware.OwnershipGuard(), h.SaveAddresses)

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

	onboarding.Put("/profile/:employeeId/identity", middleware.OwnershipGuard(), h.SaveIdentity)
	onboarding.Get("/profile/:employeeId/identity", middleware.OwnershipGuard(), h.GetIdentity)

	onboarding.Post("/profile/:employeeId/documents", middleware.OwnershipGuard(), h.UploadDocument)
	onboarding.Get("/profile/:employeeId/documents", middleware.OwnershipGuard(), h.GetDocuments)
	onboarding.Delete("/documents/:id", middleware.OwnershipGuard(func(c *fiber.Ctx) (int, error) {
		id, _ := strconv.Atoi(c.Params("id"))
		return repo.GetDocumentOwner(id)
	}), h.DeleteDocument)
	onboarding.Patch("/documents/:id/verify", hrOnly, h.VerifyDocument)

	onboarding.Post("/profile/:employeeId/assets", hrOnly, h.AssignAsset)
	onboarding.Get("/profile/:employeeId/assets", middleware.OwnershipGuard(), h.GetAssets)
	onboarding.Patch("/assets/:id/acknowledge", middleware.OwnershipGuard(func(c *fiber.Ctx) (int, error) {
		id, _ := strconv.Atoi(c.Params("id"))
		return repo.GetAssetOwner(id)
	}), h.AcknowledgeAsset)

	onboarding.Get("/profile/:employeeId/completion", middleware.OwnershipGuard(), h.GetCompletion)
	onboarding.Get("/admin/dashboard", hrOnly, h.GetDashboard)
}
