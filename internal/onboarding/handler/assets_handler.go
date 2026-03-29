package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"hrms/internal/onboarding/model"
)


func (h *OnboardingHandler) GetAssets(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Query("employee_id"))

	data, err := h.Service.GetAssets(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}


func (h *OnboardingHandler) AcknowledgeAsset(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	err := h.Service.AcknowledgeAsset(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "asset acknowledged",
	})
}


func (h *OnboardingHandler) AssignAsset(c *fiber.Ctx) error {

	var req model.EmployeeAsset

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.Service.AssignAsset(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "asset assigned",
	})
}