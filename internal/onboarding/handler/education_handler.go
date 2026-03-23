package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"hrms/internal/onboarding/model"
)

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

func (h *OnboardingHandler) UpdateEducation(c *fiber.Ctx) error {

	idParam := c.Params("id")

	id, _ := strconv.Atoi(idParam)

	var edu model.Education

	if err := c.BodyParser(&edu); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.Service.UpdateEducation(id, edu)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "education updated",
	})
}