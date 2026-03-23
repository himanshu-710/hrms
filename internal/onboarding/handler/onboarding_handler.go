package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

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

func (h *OnboardingHandler) Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Onboarding service working",
	})
}

func (h *OnboardingHandler) CreateEmployee(c *fiber.Ctx) error {

	var req struct {
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		Email      string `json:"personal_email"`
		Department string `json:"department"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := h.Service.CreateEmployee(
		req.FirstName,
		req.LastName,
		req.Email,
		req.Department,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Employee created successfully",
	})
}

func (h *OnboardingHandler) GetProfile(c *fiber.Ctx) error {

	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	emp, err := h.Service.GetEmployee(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(emp)
}

func (h *OnboardingHandler) AddEducation(c *fiber.Ctx) error {

	var edu model.Education

	if err := c.BodyParser(&edu); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.Service.AddEducation(edu)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "education added",
	})
}

func (h *OnboardingHandler) GetEducation(c *fiber.Ctx) error {

	idParam := c.Params("employeeId")

	id, _ := strconv.Atoi(idParam)

	data, err := h.Service.GetEducation(id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}

func (h *OnboardingHandler) DeleteEducation(c *fiber.Ctx) error {

	idParam := c.Params("id")

	id, _ := strconv.Atoi(idParam)

	err := h.Service.DeleteEducation(id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "education deleted",
	})
}

func (h *OnboardingHandler) AddExperience(c *fiber.Ctx) error {

	var exp model.Experience

	if err := c.BodyParser(&exp); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.Service.AddExperience(exp)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "experience added",
	})
}

func (h *OnboardingHandler) GetExperience(c *fiber.Ctx) error {

	idParam := c.Params("employeeId")

	id, _ := strconv.Atoi(idParam)

	data, err := h.Service.GetExperience(id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}

func (h *OnboardingHandler) DeleteExperience(c *fiber.Ctx) error {

	idParam := c.Params("id")

	id, _ := strconv.Atoi(idParam)

	err := h.Service.DeleteExperience(id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "experience deleted",
	})
}

func (h *OnboardingHandler) SaveAddresses(c *fiber.Ctx) error {

	var req model.AddressesRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.Service.SaveAddresses(req)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "addresses saved",
	})
}
