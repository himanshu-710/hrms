package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"hrms/internal/onboarding/model"
)

func (h *OnboardingHandler) Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Onboarding service working"})
}

func (h *OnboardingHandler) CreateEmployee(c *fiber.Ctx) error {
	var req model.CreateEmployeeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	err := h.Service.CreateEmployee(req.FirstName, req.LastName, req.Email, req.Department, req.EmploymentContextRole)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Employee created successfully"})
}

func (h *OnboardingHandler) GetProfile(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	profile, err := h.Service.GetFullProfile(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(profile)
}

func (h *OnboardingHandler) UpdatePrimaryDetails(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("employeeId"))
	if err != nil || id == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid employee id"})
	}
	var req model.PrimaryDetailsRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.Service.UpdatePrimaryDetails(id, req); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "primary details updated"})
}

func (h *OnboardingHandler) UpdateContactDetails(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("employeeId"))
	if err != nil || id == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid employee id"})
	}
	var req model.ContactRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.Service.UpdateContactDetails(id, req); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "contact details updated"})
}

func (h *OnboardingHandler) UpdateRelations(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("employeeId"))
	if err != nil || id == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid employee id"})
	}
	var req model.RelationsRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	relations := map[string]interface{}{
		"mother":   req.Mother,
		"father":   req.Father,
		"spouse":   req.Spouse,
		"children": req.Children,
	}
	if err := h.Service.UpdateRelations(id, relations); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "relations updated"})
}
