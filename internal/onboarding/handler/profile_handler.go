package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"hrms/internal/onboarding/model"

)


// ================= HEALTH =================

func (h *OnboardingHandler) Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Onboarding service working",
	})
}

// ================= CREATE EMPLOYEE =================

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

// ================= GET FULL PROFILE =================

func (h *OnboardingHandler) GetProfile(c *fiber.Ctx) error {

	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	profile, err := h.Service.GetFullProfile(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(profile)
}

// ================= PRIMARY DETAILS =================

func (h *OnboardingHandler) UpdatePrimaryDetails(c *fiber.Ctx) error {

	var req model.PrimaryDetailsRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	idParam := c.Query("employee_id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid employee_id"})
	}

	err = h.Service.UpdatePrimaryDetails(id, req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "primary details updated",
	})
}

// ================= CONTACT DETAILS =================

func (h *OnboardingHandler) UpdateContactDetails(c *fiber.Ctx) error {

	var req model.ContactRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	idParam := c.Query("employee_id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid employee_id"})
	}

	err = h.Service.UpdateContactDetails(id, req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "contact details updated",
	})
}

// ================= RELATIONS =================

func (h *OnboardingHandler) UpdateRelations(c *fiber.Ctx) error {

	var req model.RelationsRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	idParam := c.Query("employee_id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid employee_id"})
	}

	relations := map[string]interface{}{
		"mother":   req.Mother,
		"father":   req.Father,
		"spouse":   req.Spouse,
		"children": req.Children,
	}

	err = h.Service.UpdateRelations(id, relations)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "relations updated",
	})
}