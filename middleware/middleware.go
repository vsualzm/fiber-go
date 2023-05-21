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

	_, err := utils.VerifyToken(token)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unatheticated",
		})
	}

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
