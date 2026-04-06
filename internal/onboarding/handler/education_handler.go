package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"hrms/internal/onboarding/model"
	"hrms/pkg/middleware"
)

func (h *OnboardingHandler) AddEducation(c *fiber.Ctx) error {
	claims, ok := middleware.GetClaims(c)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "missing authentication claims"})
	}

	var req model.EducationRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if claims.Role != "HR" && claims.Role != "HR_ADMIN" && req.EmployeeID != claims.EmployeeID {
		return c.Status(403).JSON(fiber.Map{"error": "access denied"})
	}

	err := h.Service.AddEducation(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "education added"})
}

func (h *OnboardingHandler) GetEducation(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("employeeId"))
	if err != nil || id == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid employee id"})
	}

	data, err := h.Service.GetEducation(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}

func (h *OnboardingHandler) DeleteEducation(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid education id"})
	}

	err = h.Service.DeleteEducation(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "education deleted"})
}

func (h *OnboardingHandler) UpdateEducation(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid education id"})
	}

	var req model.EducationRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err = h.Service.UpdateEducation(id, req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "education updated"})
}
