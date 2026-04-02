package middleware

import (
	"github.com/gofiber/fiber/v2"
	"fmt"
)


func OwnershipGuard(ownerFn ...func(*fiber.Ctx) (int, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {

		claims := GetClaims(c)

		
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
			
			paramID, parseErr := parseInt(c.Params("employeeId"))
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

func parseInt(s string) (int, error) {
	var n int
	_, err := fmt.Sscanf(s, "%d", &n)
	return n, err
}