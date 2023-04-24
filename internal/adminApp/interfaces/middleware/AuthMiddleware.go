package middleware

import (
	"fmt"
	"strings"

	"github.com/alwismt/selectify/internal/infrastructure/session"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// check for session
	s := session.NewSession(nil)
	userID := s.Get(c, "id")
	if userID == nil {
		url := strings.TrimLeft(c.OriginalURL(), "/")
		return c.Redirect(fmt.Sprintf("/controlpanel/login?redirect=%s", url), 302)
	}
	return c.Next()
}
