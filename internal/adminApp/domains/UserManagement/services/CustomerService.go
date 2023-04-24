package services

import (
	"context"
	"fmt"

	"github.com/alwismt/selectify/internal/adminApp/domains/UserManagement/entities"
	"github.com/alwismt/selectify/internal/adminApp/domains/UserManagement/repositories"
	transferobjects "github.com/alwismt/selectify/internal/adminApp/interfaces/transferObjects"
)

type CustomerService interface {
	GetCustomers(ctx context.Context, search, status, orderColumn, orderDir string, pageNumber, recordsPerPage int) ([][]string, int64, error)
	Action(ctx context.Context, data *transferobjects.Action) error
}

type customerService struct {
	customerRepository repositories.CustomerRepository
}

func NewCustomerService(customerRepository repositories.CustomerRepository) CustomerService {
	return &customerService{customerRepository: customerRepository}
}

func (s *customerService) GetCustomers(ctx context.Context, search, status, orderColumn, orderDir string, pageNumber, recordsPerPage int) ([][]string, int64, error) {

	customers, count, err := s.customerRepository.GetCustomers(ctx, search, status, orderColumn, orderDir, pageNumber, recordsPerPage)
	if err != nil {
		return nil, count, err
	}

	var rows [][]string
	for _, v := range customers {
		status, filter_ban_active := s.getStatusAndAction(v)
		row := s.getRow(v, status, filter_ban_active)
		rows = append(rows, row)
	}
	return rows, count, nil
}

func (s *customerService) Action(ctx context.Context, data *transferobjects.Action) error {
	// delete customer
	if data.Action == "delete" {
		return s.customerRepository.Delete(ctx, data.ID)
	}
	var status int // 1 = active, 2 = banned
	switch data.Action {
	case "ban":
		status = 2
	case "active":
		status = 1
	}
	// update customer status
	if status == 1 || status == 2 {
		return s.customerRepository.UpdateStatus(ctx, data.ID, status)
	}
	// invalid action
	return fmt.Errorf("invalid action")
}

func (s *customerService) getStatusAndAction(customer entities.Customer) (string, string) {

	var status string
	var filter_ban_active string

	switch customer.Status {
	case 0:
		status = "Inactive"
	case 1:
		status = "Active"
		filter_ban_active = `{"name": "Ban", "url": "#", "filter_ban": "ban_row", "id": "` + customer.ID.String() + `"},`
	case 2:
		status = "Banned"
		filter_ban_active = `{"name": "Active", "url": "#", "filter_active": "active_row", "id": "` + customer.ID.String() + `"},`
	default:
		status = "Blocked"
	}

	return status, filter_ban_active
}

func (s *customerService) getRow(customer entities.Customer, status, filter_ban_active string) []string {

	return []string{
		customer.Name,
		customer.Email,
		customer.Phone,
		status,
		customer.CreatedAt.Format("2006-01-02 - 15:04"),
		`{"action":[
			{"name": "View", "url": "` + "/customer/" + customer.ID.String() + `", "target": "` + customer.ID.String() + `"}, 
			` + filter_ban_active + `
			{"name": "Delete", "url": "#", "filter_delete": "delete_row", "id": "` + customer.ID.String() + `"}
			]}`,
	}
}
