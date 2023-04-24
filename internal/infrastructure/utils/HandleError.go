package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func HandleError(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		if code == fiber.StatusNotFound {
			e.Message = "Page not found"
		}
	}
	if c.Get("X-Requested-With") == "XMLHttpRequest" {
		return c.Status(code).JSON(fiber.Map{
			"message": e.Message,
		})
	}

	err = c.Status(code).Render("errors/error", fiber.Map{
		"Title": e.Message,
		"Code":  code,
	}, "layouts/error")
	if err != nil {
		// In case the Render fails
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	// Return from handler
	return nil
}
