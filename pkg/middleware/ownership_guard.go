package middleware

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func OwnershipGuard(ownerFn ...func(*fiber.Ctx) (int, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := GetClaims(c)
		if !ok {
			return c.Status(401).JSON(fiber.Map{"error": "missing authentication claims"})
		}

		if claims.Role == "HR" || claims.Role == "HR_ADMIN" {
			return c.Next()
		}

		var ownerID int
		var err error

		if len(ownerFn) > 0 && ownerFn[0] != nil {
			ownerID, err = ownerFn[0](c)
			if err != nil {
				return c.Status(404).JSON(fiber.Map{"error": "resource not found"})
			}
		} else {
			paramID, parseErr := strconv.Atoi(c.Params("employeeId"))
			if parseErr != nil || paramID == 0 {
				return c.Status(400).JSON(fiber.Map{"error": "invalid employee id"})
			}
			ownerID = paramID
		}

		if claims.EmployeeID != ownerID {
			return c.Status(403).JSON(fiber.Map{"error": "access denied"})
		}

		return c.Next()
	}
}

func RequireRoles(roles ...string) fiber.Handler {
	allowed := make(map[string]struct{}, len(roles))
	for _, role := range roles {
		allowed[role] = struct{}{}
	}

	return func(c *fiber.Ctx) error {
		claims, ok := GetClaims(c)
		if !ok {
			return c.Status(401).JSON(fiber.Map{"error": "missing authentication claims"})
		}

		if _, ok := allowed[claims.Role]; !ok {
			return c.Status(403).JSON(fiber.Map{"error": "access denied"})
		}

		return c.Next()
	}
}
