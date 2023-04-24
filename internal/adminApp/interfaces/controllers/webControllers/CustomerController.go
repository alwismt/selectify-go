package webcontrollers

import (
	"context"
	"fmt"
	"time"

	"github.com/alwismt/selectify/internal/adminApp/domains/UserManagement/services"
	"github.com/gofiber/fiber/v2"
)

type CustomerController interface {
	Index(c *fiber.Ctx) error
}

type customerController struct {
	customerService services.CustomerService
}

func NewCustomerController(customerService services.CustomerService) CustomerController {
	return &customerController{customerService: customerService}
}

func (cc *customerController) Index(c *fiber.Ctx) error {
	start := time.Now()
	_, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	breadcrumb := fiber.Map{
		"name": "Customer Management",
	}

	cardHeader := fiber.Map{
		"search": "Search Customers",
		"cardToolbar": fiber.Map{
			"selectName": "Status",
			"values": fiber.Map{
				"active":   "Active",
				"inactive": "Inactive",
				"banned":   "Banned",
				"blocked":  "Blocked",
			},
			"button": fiber.Map{
				"url":  "#",
				"name": "Add Customer",
			},
		},
	}

	// customers, err := cc.customerService.GetCustomers(ctx)
	// if err != nil {
	// 	return err
	// }

	// return c.JSON(customers)

	err := c.Render("indexDataTable",
		fiber.Map{"title": "Customer Profiles",
			"breadcrumb":  breadcrumb,
			"cardHeader":  cardHeader,
			"active":      "Customer Management",
			"subactive":   "Customer Profiles",
			"columnNames": []string{"Name", "Email", "Phone", "Status", "Registered At"},
		},
	)
	elapsed := time.Since(start)
	fmt.Printf("Execution time1: %s\n", elapsed)
	return err
}
