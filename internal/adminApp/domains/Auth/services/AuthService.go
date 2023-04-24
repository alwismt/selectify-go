package services

import (
	"context"
	"fmt"
	"time"

	"github.com/alwismt/selectify/internal/adminApp/domains/Auth/repositories"
	transferobjects "github.com/alwismt/selectify/internal/adminApp/interfaces/transferObjects"
	"github.com/alwismt/selectify/internal/infrastructure/messagebroker/queue"
	"github.com/alwismt/selectify/internal/infrastructure/utils"
	"github.com/google/uuid"
)

type AuthService interface {
	AuthCheck(ctx context.Context, data *transferobjects.SignInDTO, eventData *transferobjects.EventAuthDTO) (uuid.UUID, string, int, error)
}

type customerService struct {
	authRepo repositories.AuthRepository
	event    queue.RabbitMQPublisher
}

func NewAuthService(authRepository repositories.AuthRepository) AuthService {
	if authRepository == nil {
		authRepository = repositories.NewAuthRepository(nil)
	}
	event := queue.NewRabbitMQPublisher(nil)
	return &customerService{authRepo: authRepository, event: event}
}

func (c *customerService) AuthCheck(ctx context.Context, data *transferobjects.SignInDTO, eventData *transferobjects.EventAuthDTO) (uuid.UUID, string, int, error) {
	user, err := c.authRepo.FindByEmail(ctx, data.Email) // Find user by email using repository
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return user.ID, user.Name, user.Type, fmt.Errorf("email or password is incorrect")
		} else {
			return user.ID, user.Name, user.Type, fmt.Errorf("something went wrong! please try again")
		}
	}
	comparePassword := utils.ComparePasswords(user.Password, data.Password) // Compare password
	if !comparePassword {
		return user.ID, user.Name, user.Type, fmt.Errorf("email or password is incorrect")
	}

	eventData.Name = "selectify_auth"   // Set event name to user
	eventData.Type = "login"            // Set event type to login
	eventData.UserID = user.ID.String() // Set user id to event data
	eventData.Timestamp = time.Now()    // Set timestamp to event data
	// err = c.event.AuthEvent(eventData)  // Publish user to message broker
	err = c.event.QueueEvent(eventData) // Publish user to message broker

	if err != nil {
		return user.ID, user.Name, user.Type, fmt.Errorf("something went wrong! please try again")
	}

	return user.ID, user.Name, user.Type, nil // Return user id and name
}

func (c *customerService) SuspiciousDetect(ctx context.Context) {

}
