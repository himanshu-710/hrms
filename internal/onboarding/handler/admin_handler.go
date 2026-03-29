package handler

import "github.com/gofiber/fiber/v2"

func (h *OnboardingHandler) GetDashboard(c *fiber.Ctx) error {

	data, err := h.Service.GetDashboard()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(data)
}