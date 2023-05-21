package middleware

import (
	"fiber-go/utils"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {

	token := c.Get("x-token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unatheticated",
		})
	}

	// _, err := utils.VerifyToken(token)
	claims, err := utils.DecodeToken(token)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unatheticated",
		})
	}

	role := claims["role"].(string)
	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Forbiden Access",
		})
	}

	// untuk logger aja
	c.Locals("userInfo", claims)
	// c.Locals("role", claims["role"])

	// ketika pengen masukin header biasa aja dengan token ini
	// if token != "secret" {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"message": "unatheticated",
	// 	})
	// }

	return c.Next()
}

func PermissionCreate(c *fiber.Ctx) error {
	return c.Next()
}
