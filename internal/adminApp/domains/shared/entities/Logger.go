package entities

import (
	"time"

	"github.com/google/uuid"
)

type Logger struct {
	ID        uuid.UUID `bson:"_id"`
	UserID    uuid.UUID `bson:"user_id"`
	SessionID uuid.UUID `bson:"session_id"`
	Message   string    `bson:"message"`   //  a description of the activity or event that occurred
	Source    string    `bson:"source"`    // the source of the log entry (e.g. module, function, line number, etc.)
	Action    string    `bson:"action"`    //  the specific action or operation that was performed (e.g. CREATE, UPDATE, DELETE, LOGIN, LOGOUT, etc.)
	Timestamp time.Time `bson:"timestamp"` // the date and time of the activity or event
}
