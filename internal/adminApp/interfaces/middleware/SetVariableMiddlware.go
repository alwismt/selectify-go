package middleware

import (
	"github.com/alwismt/selectify/internal/adminApp/interfaces/views"
	"github.com/gofiber/fiber/v2"
)

func SetVariableMiddleware(c *fiber.Ctx) error {
	sideMenu := views.GetMenuItems()
	c.Locals("sideMenuMap", sideMenu)
	return c.Next() // return next after setup local variables
}
