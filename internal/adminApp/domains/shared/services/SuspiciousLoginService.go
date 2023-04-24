package services

import (
	"context"

	"github.com/alwismt/selectify/internal/adminApp/domains/shared/entities"
	authLogRepo "github.com/alwismt/selectify/internal/adminApp/domains/shared/repositories"
	transferobjects "github.com/alwismt/selectify/internal/adminApp/interfaces/transferObjects"

	// "github.com/alwismt/selectify/internal/infrastructure/messagebroker"

	"github.com/alwismt/selectify/internal/infrastructure/messagebroker/queue"
	mongodb "github.com/alwismt/selectify/internal/infrastructure/persistence/mongoDB"
	"go.mongodb.org/mongo-driver/mongo"
)

type SuspiciousLoginService interface {
	NewSuspiciousLogin(ctx context.Context, log *entities.AuthLog) error
}

type suspiciousLoginService struct {
	db    *mongo.Database
	repo  authLogRepo.SuspiciousLoginRepository
	event queue.RabbitMQPublisher
}

func NewSuspiciousLoginService(mdb *mongo.Database) SuspiciousLoginService {
	if mdb == nil {
		mdb = mongodb.MongoDB
	}
	event := queue.NewRabbitMQPublisher(nil)
	return &suspiciousLoginService{db: mdb, repo: authLogRepo.NewSuspiciousLoginRepository(mdb), event: event}
}

func (q *suspiciousLoginService) NewSuspiciousLogin(ctx context.Context, log *entities.AuthLog) error {
	err := q.repo.Add(ctx, log)
	if err != nil {
		return err
	}

	event := &transferobjects.QueueDTO{
		Name: "selectify_event",
		Type: "suspiciousLogin",
		Data: log,
	}
	err = q.event.QueueEvent(event)
	if err != nil {
		return err
	}

	email := &transferobjects.EmailDTO{
		Name: "selectify_email",
		Type: "suspiciousLogin",
		Data: log,
	}

	err = q.event.QueueEvent(email)
	if err != nil {
		return err
	}

	return nil
}
