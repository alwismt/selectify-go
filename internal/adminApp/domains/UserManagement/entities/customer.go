package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID       uuid.UUID `gorm:"primarykey" json:"id"`
	Name     string    `gorm:"size:255;not null" json:"name"`
	Email    string    `gorm:"unique;size:255;not null" json:"email"`
	Phone    string    `gorm:"unique;size:255;not null" json:"phone"`
	Password string    `gorm:"size:255;not null" json:"-"`
	Status   int       `gorm:"not null; default:1;" json:"-"` // 0: inactive, 1: active, 2: banned, 3: blocked
	// Address  []CustomerAddress `gorm:"foreignKey:CustomerID;" json:"address"`
	// LastLogin time.Time `json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
