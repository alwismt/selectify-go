package web

import (
	"github.com/alwismt/selectify/internal/adminApp/domains/Auth/services"
	webcontrollers "github.com/alwismt/selectify/internal/adminApp/interfaces/controllers/webControllers"
	"github.com/alwismt/selectify/internal/adminApp/interfaces/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (Wb *Web) AuthRoutes(a *fiber.App) {
	// set up auth controller
	ctrl := webcontrollers.NewAuthController(services.NewAuthService(nil), nil)
	// set up customer routes
	routeGuest := a.Group("/controlpanel")
	// routeGuest.Use(sharedservices.LoggerService)

	routeGuest.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/New_York",
	}))
	routeGuest.Get("/login", middleware.GuestMiddleware, ctrl.Login)
	routeGuest.Post("/login", middleware.GuestMiddleware, ctrl.AuthCheck)

	routeAuth := a.Group("/controlpanel")
	// routeAuth.Get("/dashboard", middleware.Auth, ctrl.Login)
	routeAuth.Get("/logout", middleware.AuthMiddleware, ctrl.Logout)
}
