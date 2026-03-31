package handler

import (
    "strconv"
    "github.com/gofiber/fiber/v2"
    "hrms/internal/onboarding/model"
)

func (h *OnboardingHandler) GetAssets(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("employeeId"))  
    if err != nil || id == 0 {
        return c.Status(400).JSON(fiber.Map{"error": "invalid employee id"})
    }
    data, err := h.Service.GetAssets(id)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(data)
}

func (h *OnboardingHandler) AcknowledgeAsset(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    if err := h.Service.AcknowledgeAsset(id); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"message": "asset acknowledged"})
}

func (h *OnboardingHandler) AssignAsset(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("employeeId"))
    if err != nil || id == 0 {
        return c.Status(400).JSON(fiber.Map{"error": "invalid employee id"})
    }
    var req model.AssignAssetRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }
    req.EmployeeID = id  
    if err := h.Service.AssignAsset(req); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"message": "asset assigned"})
}