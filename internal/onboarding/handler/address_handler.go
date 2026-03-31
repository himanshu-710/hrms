package handler

import (
    "strconv"
    "github.com/gofiber/fiber/v2"
    "hrms/internal/onboarding/model"
)

func (h *OnboardingHandler) SaveAddresses(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("employeeId"))  
    if err != nil || id == 0 {
        return c.Status(400).JSON(fiber.Map{"error": "invalid employee id"})
    }
    var req model.AddressesRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }
    req.EmployeeID = id  
    if err := h.Service.SaveAddresses(req); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"message": "addresses saved"})
}