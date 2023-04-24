package repositories

import (
	"context"

	"github.com/alwismt/selectify/internal/adminApp/domains/Auth/entities"
	mongodb "github.com/alwismt/selectify/internal/infrastructure/persistence/mongoDB"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository interface {
	FindByEmail(ctx context.Context, email string) (entities.Admin, error)
}

type authRepository struct {
	db *mongo.Database
}

func NewAuthRepository(mdb *mongo.Database) AuthRepository {
	if mdb == nil {
		mdb = mongodb.MongoDB
	}
	return &authRepository{db: mdb}
}

func (r *authRepository) FindByEmail(ctx context.Context, email string) (entities.Admin, error) {
	admin := entities.Admin{}
	filter := bson.M{"email": email}
	err := r.db.Collection("admins").FindOne(ctx, filter).Decode(&admin)
	if err != nil {
		return admin, err
	}
	return admin, nil
}
