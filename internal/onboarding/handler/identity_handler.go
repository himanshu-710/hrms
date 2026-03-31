package handler

import (
    "strconv"
    "github.com/gofiber/fiber/v2"
    "hrms/internal/onboarding/model"
)

func (h *OnboardingHandler) SaveIdentity(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("employeeId"))
    if err != nil || id == 0 {
        return c.Status(400).JSON(fiber.Map{"error": "invalid employee id"})
    }
    var req model.IdentityRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }
    req.EmployeeID = id  
    if err := h.Service.SaveIdentity(req); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"message": "identity saved"})
}

func (h *OnboardingHandler) GetIdentity(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("employeeId")) 
    if err != nil || id == 0 {
        return c.Status(400).JSON(fiber.Map{"error": "invalid employee_id"})
    }
    data, err := h.Service.GetIdentity(id)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(data)
}