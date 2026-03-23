package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"hrms/internal/onboarding/model"
)

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

func (h *OnboardingHandler) UpdateExperience(c *fiber.Ctx) error {

	idParam := c.Params("id")

	id, _ := strconv.Atoi(idParam)

	var exp model.Experience

	if err := c.BodyParser(&exp); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.Service.UpdateExperience(id, exp)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "experience updated",
	})
}