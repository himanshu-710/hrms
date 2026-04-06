package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"hrms/internal/onboarding/model"
	"hrms/pkg/middleware"
)

func (h *OnboardingHandler) AddExperience(c *fiber.Ctx) error {
	claims, ok := middleware.GetClaims(c)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "missing authentication claims"})
	}

	var req model.ExperienceRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if claims.Role != "HR" && claims.Role != "HR_ADMIN" && req.EmployeeID != claims.EmployeeID {
		return c.Status(403).JSON(fiber.Map{"error": "access denied"})
	}

	err := h.Service.AddExperience(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "experience added"})
}

func (h *OnboardingHandler) GetExperience(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("employeeId"))
	if err != nil || id == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid employee id"})
	}

	data, err := h.Service.GetExperience(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}

func (h *OnboardingHandler) DeleteExperience(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid experience id"})
	}

	err = h.Service.DeleteExperience(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "experience deleted"})
}

func (h *OnboardingHandler) UpdateExperience(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid experience id"})
	}

	var req model.ExperienceRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err = h.Service.UpdateExperience(id, req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "experience updated"})
}
