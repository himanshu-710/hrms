package routes

import (
	"github.com/gin-gonic/gin"

	"hrms/internal/onboarding/handler"
	"hrms/internal/onboarding/repository"
	"hrms/internal/onboarding/service"
	"hrms/pkg/database"
)

func RegisterOnboardingRoutes(r *gin.RouterGroup) {

	repo := repository.NewOnboardingRepository(database.DB)
	service := service.NewOnboardingService(repo)
	handler := handler.NewOnboardingHandler(service)

	onboarding := r.Group("/onboarding")

	onboarding.GET("/health", handler.Health)
	onboarding.POST("/employee", handler.CreateEmployee)

	onboarding.POST("/profile", handler.CreateEmployee)
	onboarding.GET("/profile/:id", handler.GetProfile)
	onboarding.POST("/education", handler.AddEducation)

	onboarding.GET("/education/:employeeId", handler.GetEducation)

	onboarding.DELETE("/education/:id", handler.DeleteEducation)

	onboarding.POST("/experience", handler.AddExperience)

	onboarding.GET("/experience/:employeeId", handler.GetExperience)

	onboarding.DELETE("/experience/:id", handler.DeleteExperience)

	onboarding.PUT("/addresses", handler.SaveAddresses)
}
