package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"hrms/pkg/middleware"
)

func RegisterAuthRoutes(app *fiber.App, h interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}) {
	// Rate limiter — 10 requests per minute per IP (token bucket)
	rateLimiter := limiter.New(limiter.Config{
		Max:        10,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).JSON(fiber.Map{"error": "too many requests, slow down"})
		},
	})

	auth := app.Group("/api/v1/auth", rateLimiter)

	auth.Post("/register", h.Register)
	auth.Post("/login", h.Login)
	auth.Post("/refresh", h.RefreshToken)
	auth.Post("/logout", middleware.AuthMiddleware(), h.Logout)
}