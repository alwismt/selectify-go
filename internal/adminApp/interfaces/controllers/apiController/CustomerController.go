package apicontroller

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/alwismt/selectify/internal/adminApp/domains/UserManagement/services"
	transferobjects "github.com/alwismt/selectify/internal/adminApp/interfaces/transferObjects"
	"github.com/alwismt/selectify/internal/infrastructure/utils"
	"github.com/gofiber/fiber/v2"
)

type CustomerController interface {
	DataTables(c *fiber.Ctx) error
	TableAction(c *fiber.Ctx) error
}

type customerController struct {
	customerService services.CustomerService
}

func NewCustomerController(customerService services.CustomerService) CustomerController {
	return &customerController{customerService: customerService}
}

func (cc *customerController) DataTables(c *fiber.Ctx) error {

	fmt.Println("req type", c.Get("X-Requested-With"))
	start := time.Now()
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	var (
		pageNumber     = 1
		recordsPerPage = 10
		err            = error(nil)
		search         = c.FormValue("search[value]")
		draw           = c.FormValue("draw")
		length         = c.FormValue("length")
		page           = c.FormValue("start")
		status         = c.FormValue("status")
		orderColumn    = c.FormValue("order[0][column]")
		orderDir       = c.FormValue("order[0][dir]")
	)
	if length != "" {
		recordsPerPage, err = strconv.Atoi(length) // set records per page
		if err != nil {
			return c.SendString(err.Error())
		}
	}
	if page != "" {
		startNumber, err := strconv.Atoi(page) // set start number
		if err != nil {
			return c.SendString(err.Error())
		}
		pageNumber = startNumber/recordsPerPage + 1 // set page number
	}
	// check status "all", "0", "1", "2", "3"
	switch status {
	case "":
		status = "all"
	case "inactive":
		status = "0"
	case "active":
		status = "1"
	case "banned":
		status = "2"
	case "blocked":
		status = "3"
	default:
		status = "all"
	}
	// check order column "0", "1", "2", "3", "4" (name, email, phone, status, created_at)
	switch orderColumn {
	case "0":
		orderColumn = "name"
	case "1":
		orderColumn = "email"
	case "2":
		orderColumn = "phone"
	case "3":
		orderColumn = "status"
	case "4":
		orderColumn = "created_at"
	default:
		orderColumn = "created_at"
	}
	// check order direction "asc", "desc"
	switch orderDir {
	case "asc":
		orderDir = "asc"
	case "desc":
		orderDir = "desc"
	default:
		orderDir = "asc"
	}
	// Get the paginated and filtered data from the customer service.
	customers, count, err := cc.customerService.GetCustomers(ctx, search, status, orderColumn, orderDir, pageNumber, recordsPerPage)
	if err != nil {
		return err
	}

	elapsed := time.Since(start)
	fmt.Printf("Execution time2: %s\n", elapsed)

	return c.Status(200).JSON(fiber.Map{
		"draw":            draw,
		"recordsTotal":    count,
		"recordsFiltered": count,
		"data":            customers,
	})
}

func (cc *customerController) TableAction(c *fiber.Ctx) error {
	fmt.Println("req type", c.Get("X-Requested-With"))
	data := new(transferobjects.Action) // Create a new instance of the data model
	// Parse the body into the data model
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": true,
			"msg":   "Data type mismatch",
		})
	}
	validate := utils.NewValidator() // Create a new validator instance
	if err := validate.Struct(data); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": true,
			"msg":   "ID or Action fields are missing",
		})

	}
	// Create a context with timeout
	ctx, cancel := utils.CreateContextWithTimeout()
	defer cancel() // The cancel should be deferred so resources are cleaned up

	err := cc.customerService.Action(ctx, data) // Call the service to perform the action
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.SendStatus(204) // 204 No Content
}
