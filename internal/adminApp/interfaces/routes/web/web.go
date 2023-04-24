package web

import (
	// webcontrollers "github.com/alwismt/selectify/internal/adminApp/interfaces/controllers/webControllers"
	"github.com/gofiber/fiber/v2"
)

type Web struct {
	// CustomerCtrl webcontrollers.CustomerController
}

// func CustomerController(customerCtrl webcontrollers.CustomerController) *Web {
// 	return &Web{CustomerCtrl: customerCtrl}
// }

func WebRoutes(app *fiber.App) {
	web := Web{}
	// Customer Management routes
	web.CustomerRoutes(app)
	web.AuthRoutes(app)

}
