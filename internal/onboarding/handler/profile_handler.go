package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

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

func (h *OnboardingHandler) UpdateProfile(c *fiber.Ctx) error {

	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	var req struct {
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		Email      string `json:"personal_email"`
		Department string `json:"department"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.Service.UpdateEmployee(id, req.FirstName, req.LastName, req.Email, req.Department); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	emp, err := h.Service.GetEmployee(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(emp)
}