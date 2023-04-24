package transferobjects

import "github.com/google/uuid"

type Action struct {
	ID     uuid.UUID `json:"id" validate:"required"`
	Action string    `json:"action" validate:"required"`
}
