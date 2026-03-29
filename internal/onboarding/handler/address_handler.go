package handler

import (
	"github.com/gofiber/fiber/v2"
	"hrms/internal/onboarding/model"
)

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