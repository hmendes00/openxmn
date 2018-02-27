package votes

import uuid "github.com/satori/go.uuid"

// Delete represents a delete vote on an organization transaction
type Delete struct {
	ID *uuid.UUID `json:"id"`
}

// CreateDelete creates a new Delete instance
func CreateDelete(id *uuid.UUID) *Delete {
	out := Delete{
		ID: id,
	}

	return &out
}

// GetID returns the ID
func (del *Delete) GetID() *uuid.UUID {
	return del.ID
}
