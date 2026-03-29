package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"hrms/internal/onboarding/model"
)

func (h *OnboardingHandler) AddExperience(c *fiber.Ctx) error {

	var req model.ExperienceRequest  

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.Service.AddExperience(req)  
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "experience added"})
}

func (h *OnboardingHandler) GetExperience(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("employeeId"))

	data, err := h.Service.GetExperience(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}

func (h *OnboardingHandler) DeleteExperience(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	err := h.Service.DeleteExperience(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "experience deleted"})
}

func (h *OnboardingHandler) UpdateExperience(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	var req model.ExperienceRequest  

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.Service.UpdateExperience(id, req)  
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "experience updated"})
}