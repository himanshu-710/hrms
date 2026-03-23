package utils

import "github.com/gofiber/fiber/v2"

func Success(c *fiber.Ctx, data interface{}) error {
	return c.Status(200).JSON(fiber.Map{
		"data": data,
	})
}

func Error(c *fiber.Ctx, msg string) error {
	return c.Status(500).JSON(fiber.Map{
		"error": msg,
	})
}