package middleware

import (
	"github.com/alwismt/selectify/internal/infrastructure/session"
	"github.com/gofiber/fiber/v2"
)

func SuperAdminAccessControl(c *fiber.Ctx) error {
	s := session.NewSession(nil)
	userType := s.Get(c, "type")
	if userType != nil {
		if userType == 1 {
			return c.Next()
		}
	}
	// return error forbidden
	return fiber.NewError(fiber.StatusForbidden, "Access denied | ACL Blocked")
}

func CustomerProfileACL(c *fiber.Ctx) error {
	s := session.NewSession(nil)
	userType := s.Get(c, "type")
	if userType != nil {
		if userType == 3 || userType == 2 {
			return c.Next()
		}
	}
	return fiber.NewError(fiber.StatusForbidden, "Access denied | ACL Blocked")
}
