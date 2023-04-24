package middleware

import (
	"github.com/gofiber/fiber/v2"

	"github.com/alwismt/selectify/internal/infrastructure/session"
)

func GuestMiddleware(c *fiber.Ctx) error {
	s := session.NewSession(nil)
	userID := s.Get(c, "id")
	if userID == nil {
		return c.Next()
	}
	return c.Redirect("/controlpanel/dashboard")
}
