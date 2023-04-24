package repositories

import (
	"context"

	"github.com/alwismt/selectify/internal/adminApp/domains/shared/entities"
	mongodb "github.com/alwismt/selectify/internal/infrastructure/persistence/mongoDB"
	"go.mongodb.org/mongo-driver/mongo"
)

type SuspiciousLoginRepository interface {
	Add(ctx context.Context, log *entities.AuthLog) error
}

type suspiciousLoginRepository struct {
	db *mongo.Database
}

func NewSuspiciousLoginRepository(mdb *mongo.Database) SuspiciousLoginRepository {
	if mdb == nil {
		mdb = mongodb.MongoDB
	}
	return &suspiciousLoginRepository{db: mdb}
}

func (q *suspiciousLoginRepository) Add(ctx context.Context, log *entities.AuthLog) error {
	// save to mongoDB
	_, err := q.db.Collection("suspicious_logins").InsertOne(ctx, log)
	if err != nil {
		return err
	}
	return nil
}
