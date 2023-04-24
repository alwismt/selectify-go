package web

import (
	"github.com/alwismt/selectify/internal/adminApp/domains/UserManagement/repositories"
	"github.com/alwismt/selectify/internal/adminApp/domains/UserManagement/services"
	webcontrollers "github.com/alwismt/selectify/internal/adminApp/interfaces/controllers/webControllers"
	"github.com/alwismt/selectify/internal/adminApp/interfaces/middleware"
	sqldb "github.com/alwismt/selectify/internal/infrastructure/persistence/sqlDB"
	"github.com/gofiber/fiber/v2"
)

func (Wb *Web) CustomerRoutes(a *fiber.App) {
	// set up customer controller
	ctrl := webcontrollers.NewCustomerController(services.NewCustomerService(repositories.NewCustomerRepository(sqldb.DB)))
	// set up customer routes
	route := a.Group("/controlpanel/customer")
	route.Use(middleware.AuthMiddleware)     // set up auth middleware
	route.Use(middleware.CustomerProfileACL) // set up customer profile acl middleware

	route.Get("/", middleware.SetVariableMiddleware, ctrl.Index)
}
