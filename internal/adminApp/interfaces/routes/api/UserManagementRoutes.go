package api

import (
	"github.com/alwismt/selectify/internal/adminApp/domains/UserManagement/repositories"
	"github.com/alwismt/selectify/internal/adminApp/domains/UserManagement/services"
	apicontroller "github.com/alwismt/selectify/internal/adminApp/interfaces/controllers/apiController"
	"github.com/alwismt/selectify/internal/adminApp/interfaces/middleware"
	sqldb "github.com/alwismt/selectify/internal/infrastructure/persistence/sqlDB"
	"github.com/gofiber/fiber/v2"
)

func (Wb *API) CustomerRoutes(a *fiber.App) {
	ctrl := apicontroller.NewCustomerController(services.NewCustomerService(repositories.NewCustomerRepository(sqldb.DB)))

	route := a.Group("/api/controlpanel/customer") // set up customer routes group
	route.Use(middleware.AuthMiddleware)           // set up auth middleware
	route.Use(middleware.JsonMiddleware)           // set up json middleware
	route.Use(middleware.CustomerProfileACL)       // set up customer profile acl middleware

	route.Post("/", ctrl.DataTables)        // customer datatable
	route.Post("/action", ctrl.TableAction) // customer action
}
