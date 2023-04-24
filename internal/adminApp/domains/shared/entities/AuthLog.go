package entities

import (
	"time"

	"github.com/google/uuid"
)

type AuthLog struct {
	ID              uuid.UUID `bson:"_id"`
	UserID          uuid.UUID `bson:"user_id"`
	SessionID       uuid.UUID `bson:"session_id, unique"`
	IP              string    `bson:"ip"`
	Location        string    `bson:"location"`
	Platform        string    `bson:"platform"`
	OperatingSystem string    `bson:"operating_system"`
	Model           string    `bson:"model"`
	Browser         string    `bson:"browser"`
	UserAgent       string    `bson:"user_agent"`
	ISP             string    `bson:"isp"`
	Latitude        float64   `bson:"latitude"`
	Longitude       float64   `bson:"longitude"`
	LoggedTime      time.Time `bson:"logged_time"`
	// Successful      bool      `bson:"successful"`
	CreatedAt time.Time `bson:"created_at"`
}
