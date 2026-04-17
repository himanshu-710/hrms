package handler

import (
	"context"
	"strconv"

	"hrms/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *OnboardingHandler) GetMyNotifications(c *fiber.Ctx) error {
	claims, ok := middleware.GetClaims(c)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "missing authentication claims"})
	}

	notifications, err := h.Service.GetNotifications(context.Background(), claims.EmployeeID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(notifications)
}

func (h *OnboardingHandler) MarkNotificationRead(c *fiber.Ctx) error {
	claims, ok := middleware.GetClaims(c)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "missing authentication claims"})
	}

	notificationID, err := strconv.Atoi(c.Params("id"))
	if err != nil || notificationID == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid notification id"})
	}

	if err := h.Service.MarkNotificationRead(context.Background(), claims.EmployeeID, notificationID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "notification marked as read"})
}
