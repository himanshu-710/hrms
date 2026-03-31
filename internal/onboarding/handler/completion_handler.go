package handler

import (
    "strconv"
    "github.com/gofiber/fiber/v2"
)

func (h *OnboardingHandler) GetCompletion(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("employeeId"))  
    if err != nil || id == 0 {
        return c.Status(400).JSON(fiber.Map{"error": "invalid employee_id"})
    }
    data, err := h.Service.ComputeCompletion(id)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(data)
}