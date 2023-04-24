package transferobjects

import "time"

type EventAuthDTO struct {
	IP        string    `json:"ip"`
	Name      string    `json:"name"`
	UserAgent string    `json:"user_agent"`
	SessionID string    `json:"session_id"`
	UserID    string    `json:"user_id"`
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
}
