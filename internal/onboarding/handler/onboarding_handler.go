package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"hrms/internal/onboarding/model"
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
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		Email      string `json:"personal_email"`
		Department string `json:"department"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.Service.CreateEmployee(
		req.FirstName,
		req.LastName,
		req.Email,
		req.Department,
	)

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
func (h *OnboardingHandler) GetProfile(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	emp, err := h.Service.GetEmployee(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, emp)
}

func (h *OnboardingHandler) AddEducation(c *gin.Context) {

	var edu model.Education

	if err := c.ShouldBindJSON(&edu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.AddEducation(edu)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "education added",
	})
}
func (h *OnboardingHandler) GetEducation(c *gin.Context) {

	idParam := c.Param("employeeId")

	id, _ := strconv.Atoi(idParam)

	data, err := h.Service.GetEducation(id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func (h *OnboardingHandler) DeleteEducation(c *gin.Context) {

	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	err := h.Service.DeleteEducation(id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "education deleted",
	})
}

func (h *OnboardingHandler) AddExperience(c *gin.Context) {

	var exp model.Experience

	if err := c.ShouldBindJSON(&exp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.AddExperience(exp)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "experience added",
	})
}
func (h *OnboardingHandler) GetExperience(c *gin.Context) {

	idParam := c.Param("employeeId")

	id, _ := strconv.Atoi(idParam)

	data, err := h.Service.GetExperience(id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func (h *OnboardingHandler) DeleteExperience(c *gin.Context) {

	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	err := h.Service.DeleteExperience(id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "experience deleted",
	})
}

func (h *OnboardingHandler) SaveAddresses(c *gin.Context) {

	var req model.AddressesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.SaveAddresses(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "addresses saved",
	})
}
