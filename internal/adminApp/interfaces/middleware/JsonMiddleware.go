package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func JsonMiddleware(c *fiber.Ctx) error {
	if c.Get("X-Requested-With") == "XMLHttpRequest" {
		return c.Next()
	}
	// return error forbidden
	err := fiber.NewError(fiber.StatusForbidden, "Method not allowed | XHR request only")
	return err
}
