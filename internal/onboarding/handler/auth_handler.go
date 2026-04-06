package handler

import (
	"hrms/internal/onboarding/model"
	"hrms/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *OnboardingHandler) Register(c *fiber.Ctx) error {
	var req model.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if req.WorkEmail == "" || req.Password == "" || req.EmployeeCode == "" {
		return c.Status(400).JSON(fiber.Map{"error": "work_email, password and employee_code are required"})
	}
	if err := h.Service.Register(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "registered successfully"})
}

func (h *OnboardingHandler) Login(c *fiber.Ctx) error {
	var req model.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if req.WorkEmail == "" || req.Password == "" {
		return c.Status(400).JSON(fiber.Map{"error": "work_email and password are required"})
	}
	resp, err := h.Service.Login(req)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}

func (h *OnboardingHandler) RefreshToken(c *fiber.Ctx) error {
	var req model.RefreshRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if req.RefreshToken == "" {
		return c.Status(400).JSON(fiber.Map{"error": "refresh_token is required"})
	}
	resp, err := h.Service.Refresh(req)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}

func (h *OnboardingHandler) Logout(c *fiber.Ctx) error {
	claims, ok := middleware.GetClaims(c)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "missing authentication claims"})
	}
	if err := h.Service.Logout(claims.EmployeeID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "logged out successfully"})
}
