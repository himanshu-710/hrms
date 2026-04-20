package handler

import (
	"hrms/internal/onboarding/model"

	"github.com/gofiber/fiber/v2"
)

func (h *OnboardingHandler) GetDashboard(c *fiber.Ctx) error {

	data, err := h.Service.GetDashboard()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(data)
}

func (h *OnboardingHandler) CreateEmployee(c *fiber.Ctx) error {
	var req model.CreateEmployeeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	emp, err := h.Service.CreateEmployee(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(emp)
}
