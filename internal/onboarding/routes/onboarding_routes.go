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
}