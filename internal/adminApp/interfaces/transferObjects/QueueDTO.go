package transferobjects

import "time"

type QueueDTO struct {
	Name      string      `json:"name"`
	Type      string      `json:"type"`
	Data      interface{} `json:"data"`
	Timestamp time.Time   `json:"timestamp"`
}
