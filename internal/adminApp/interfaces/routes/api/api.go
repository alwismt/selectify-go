package api

import "github.com/gofiber/fiber/v2"

type API struct{}

func APIRoutes(app *fiber.App) {
	api := API{}

	// Customer Management routes
	api.CustomerRoutes(app)

}
