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
    store, err := storage.NewMinIOStorage()
if err != nil {
    panic(err)
}
    svc := service.NewOnboardingService(repo, store)
    h := handler.NewOnboardingHandler(svc)

    onboarding := app.Group("/api/v1/onboarding")

    onboarding.Get("/health", h.Health)

    
    onboarding.Post("/employee", h.CreateEmployee)
    onboarding.Get("/profile/:id", h.GetProfile)
    onboarding.Put("/profile/:employeeId/primary", h.UpdatePrimaryDetails)
    onboarding.Put("/profile/:employeeId/contact", h.UpdateContactDetails)
    onboarding.Put("/profile/:employeeId/relations", h.UpdateRelations)
    onboarding.Put("/profile/:employeeId/addresses", h.SaveAddresses)

  
    onboarding.Post("/education", h.AddEducation)
    onboarding.Get("/education/:employeeId", h.GetEducation)
    onboarding.Put("/education/:id", h.UpdateEducation)
    onboarding.Delete("/education/:id", h.DeleteEducation)

    
    onboarding.Post("/experience", h.AddExperience)
    onboarding.Get("/experience/:employeeId", h.GetExperience)
    onboarding.Put("/experience/:id", h.UpdateExperience)
    onboarding.Delete("/experience/:id", h.DeleteExperience)

   
    onboarding.Put("/profile/:employeeId/identity", h.SaveIdentity)
    onboarding.Get("/profile/:employeeId/identity", h.GetIdentity)

    
    onboarding.Post("/profile/:employeeId/documents", h.UploadDocument)
    onboarding.Get("/profile/:employeeId/documents", h.GetDocuments)
    onboarding.Delete("/documents/:id", h.DeleteDocument)
    onboarding.Patch("/documents/:id/verify", h.VerifyDocument)

    
    onboarding.Post("/profile/:employeeId/assets", h.AssignAsset)
    onboarding.Get("/profile/:employeeId/assets", h.GetAssets)
    onboarding.Patch("/assets/:id/acknowledge", h.AcknowledgeAsset)

   
    onboarding.Get("/profile/:employeeId/completion", h.GetCompletion)
    onboarding.Get("/admin/dashboard", h.GetDashboard)
}