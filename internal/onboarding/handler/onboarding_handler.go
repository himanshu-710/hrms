package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hrms/internal/onboarding/service"
)

type OnboardingHandler struct {
	Service *service.OnboardingService
}

func NewOnboardingHandler(service *service.OnboardingService) *OnboardingHandler {
	return &OnboardingHandler{
		Service: service,
	}
}

func (h *OnboardingHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Onboarding service working",
	})
}

func (h *OnboardingHandler) CreateEmployee(c *gin.Context) {

	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.Service.CreateEmployee(req.Name, req.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Employee created successfully",
	})
}