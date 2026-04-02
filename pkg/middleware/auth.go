package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type ContextClaims struct {
	EmployeeID int
	Role       string
}

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(401).JSON(fiber.Map{"error": "missing or invalid authorization header"})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		secret := []byte(os.Getenv("JWT_SECRET"))

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return secret, nil
		})

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{"error": "invalid or expired token"})
		}

		mapClaims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(401).JSON(fiber.Map{"error": "invalid token claims"})
		}

		employeeID := int(mapClaims["employee_id"].(float64))
		role := mapClaims["role"].(string)

		c.Locals("claims", ContextClaims{
			EmployeeID: employeeID,
			Role:       role,
		})

		return c.Next()
	}
}


func GetClaims(c *fiber.Ctx) ContextClaims {
	return c.Locals("claims").(ContextClaims)
}