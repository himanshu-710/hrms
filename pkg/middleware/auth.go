package middleware

import (
	"fmt"
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

		employeeID, err := getIntClaim(mapClaims, "employee_id")
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "invalid employee id claim"})
		}

		role, err := getStringClaim(mapClaims, "role")
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "invalid role claim"})
		}

		c.Locals("claims", ContextClaims{
			EmployeeID: employeeID,
			Role:       role,
		})

		return c.Next()
	}
}

func getIntClaim(claims jwt.MapClaims, key string) (int, error) {
	raw, ok := claims[key]
	if !ok {
		return 0, fmt.Errorf("missing claim %s", key)
	}

	value, ok := raw.(float64)
	if !ok {
		return 0, fmt.Errorf("claim %s has invalid type", key)
	}

	return int(value), nil
}

func getStringClaim(claims jwt.MapClaims, key string) (string, error) {
	raw, ok := claims[key]
	if !ok {
		return "", fmt.Errorf("missing claim %s", key)
	}

	value, ok := raw.(string)
	if !ok || value == "" {
		return "", fmt.Errorf("claim %s has invalid type", key)
	}

	return value, nil
}

func GetClaims(c *fiber.Ctx) (ContextClaims, bool) {
	claims, ok := c.Locals("claims").(ContextClaims)
	return claims, ok
}
