package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"hrms/internal/onboarding/model"
)

func (h *OnboardingHandler) AddEducation(c *fiber.Ctx) error {

	var req model.EducationRequest  // changed from model.Education

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.Service.AddEducation(req)  // changed
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "education added"})
}

func (h *OnboardingHandler) GetEducation(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("employeeId"))

	data, err := h.Service.GetEducation(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}

func (h *OnboardingHandler) DeleteEducation(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	err := h.Service.DeleteEducation(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "education deleted"})
}

func (h *OnboardingHandler) UpdateEducation(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	var req model.EducationRequest  // changed from model.Education

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.Service.UpdateEducation(id, req)  // changed
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "education updated"})
}