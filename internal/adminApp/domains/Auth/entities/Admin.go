package entities

import (
	"time"

	"github.com/google/uuid"
)

type Admin struct {
	ID        uuid.UUID  `bson:"_id"`
	Name      string     `bson:"name"`
	Email     string     `bson:"email, unique"`
	Password  string     `bson:"password"`
	Type      int        `bson:"type"` // 1: Super Admin, 2: Admin (Default)
	Status    int        `bson:"status"`
	LastLogin time.Time  `bson:"last_login"`
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at"`
}
