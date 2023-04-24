package repositories

import (
	"context"

	"github.com/alwismt/selectify/internal/adminApp/domains/shared/entities"
	mongodb "github.com/alwismt/selectify/internal/infrastructure/persistence/mongoDB"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthLogRepository interface {
	AddAuthLog(ctx context.Context, log *entities.AuthLog) error
	GetAuthLogs(ctx context.Context, userId uuid.UUID) ([]entities.AuthLog, error)
}

type authLogRepository struct {
	db *mongo.Database
}

func NewAuthLogRepository(mdb *mongo.Database) AuthLogRepository {
	if mdb == nil {
		mdb = mongodb.MongoDB
	}
	return &authLogRepository{db: mdb}
}

func (q *authLogRepository) AddAuthLog(ctx context.Context, log *entities.AuthLog) error {
	// save to mongoDB
	_, err := q.db.Collection("auth_logs").InsertOne(ctx, log)
	if err != nil {
		return err
	}
	return nil
}

func (q *authLogRepository) GetAuthLogs(ctx context.Context, userId uuid.UUID) ([]entities.AuthLog, error) {
	// get collection auth_logs from mongoDB
	var authLogs []entities.AuthLog
	// get last 10 documents auth_logs from mongoDB using userId
	options := options.Find().SetSort(bson.M{"created_at": -1}).SetLimit(50)
	cursor, err := q.db.Collection("auth_logs").Find(ctx, bson.M{"user_id": userId}, options)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &authLogs); err != nil {
		return nil, err
	}
	return authLogs, nil
}
