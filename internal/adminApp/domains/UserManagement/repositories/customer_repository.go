package repositories

import (
	"context"
	"fmt"
	"strconv"

	"github.com/alwismt/selectify/internal/adminApp/domains/UserManagement/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	FindAll(ctx context.Context) ([]*entities.Customer, error)
	GetCustomers(ctx context.Context, search, status, orderColumn, orderDir string, pageNumber, recordsPerPage int) ([]entities.Customer, int64, error)
	Delete(ctx context.Context, id uuid.UUID) error
	UpdateStatus(ctx context.Context, id uuid.UUID, status int) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) FindAll(ctx context.Context) ([]*entities.Customer, error) {
	var customers []*entities.Customer
	r.db.WithContext(ctx).Find(&customers)
	return customers, nil
}

func (r *customerRepository) GetCustomers(ctx context.Context, search, status, orderColumn, orderDir string, pageNumber, recordsPerPage int) ([]entities.Customer, int64, error) {

	var customers []entities.Customer // slice of customers
	var count int64                   // total records
	db := r.db.WithContext(ctx)       // set db to context
	query := db.Model(customers)      // set query to db model

	// search for customer by name, email or phone
	if search != "" {
		query = query.Where("name ILIKE ? OR email ILIKE ? OR phone = ?", "%"+search+"%", "%"+search+"%", search)
	}
	// filter by status
	if status != "all" {
		st, _ := strconv.Atoi(status)
		query = query.Where("status = ?", st)
	}
	// get total records
	if err := query.Count(&count).Error; err != nil {
		return nil, count, err
	}

	order := fmt.Sprintf("%s %s", orderColumn, orderDir) // set order
	offset := (pageNumber - 1) * recordsPerPage          // set offset
	// get records
	if err := query.Order(order).Offset(offset).Limit(recordsPerPage).Find(&customers).Error; err != nil {
		return nil, count, err
	}

	return customers, count, nil
}

func (r *customerRepository) Delete(ctx context.Context, id uuid.UUID) error {
	db := r.db.WithContext(ctx) // set db to context
	// delete customer
	if err := db.Where("id = ?", id).Delete(&entities.Customer{}).Error; err != nil {
		return err // return error
	}
	return nil // return nil
}

// Status update for customer
func (r *customerRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status int) error {
	db := r.db.WithContext(ctx) // set db to context
	// update customer status
	if err := db.Model(&entities.Customer{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return err // return error
	}
	return nil // return nil
}
