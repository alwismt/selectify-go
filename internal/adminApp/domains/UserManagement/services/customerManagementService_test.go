package services_test

// func TestGetCustomers_Success(t *testing.T) {
// 	// Create mock repository with expected result
// 	expectedCustomers := []*entities.Customer{
// 		{ID: 1, FirstName: "John Doe"},
// 		{ID: 2, FirstName: "Jane Doe"},
// 	}
// 	mockRepo := repositories.NewCustomerRepository(&gorm.DB{})

// 	// Create service and call GetCustomers method
// 	customerService := services.NewCustomerService(mockRepo)
// 	ctx := context.Background()
// 	customers, err := customerService.GetCustomers(ctx)

// 	// Assert results
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedCustomers, customers)
// }

// func TestGetCustomers_Error(t *testing.T) {
// 	// Create mock repository with expected error
// 	expectedErr := errors.New("error getting customers")
// 	mockRepo := &repositories.MockCustomerRepository{
// 		FindAllFunc: func() ([]*entities.Customer, error) {
// 			return nil, expectedErr
// 		},
// 	}

// 	// Create service and call GetCustomers method
// 	customerService := services.NewCustomerService(mockRepo)
// 	ctx := context.Background()
// 	customers, err := customerService.GetCustomers(ctx)

// 	// Assert error
// 	assert.Error(t, err)
// 	assert.Nil(t, customers)
// 	assert.Equal(t, expectedErr, err)
// }
