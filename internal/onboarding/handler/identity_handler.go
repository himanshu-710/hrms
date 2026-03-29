package handler

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
	"hrms/internal/onboarding/model"
)

func (h *OnboardingHandler) SaveIdentity(c *fiber.Ctx) error {

	var req model.IdentityRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.Service.SaveIdentity(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "identity saved",
	})
}

func (h *OnboardingHandler) GetIdentity(c *fiber.Ctx) error {

	idParam := c.Query("employee_id")

	id, err := strconv.Atoi(idParam)
	if err != nil || id == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid employee_id"})
	}

	data, err := h.Service.GetIdentity(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}