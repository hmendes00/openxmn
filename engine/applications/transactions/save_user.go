package transactions

import (
	uuid "github.com/satori/go.uuid"
)

// SaveUser represents a save user transaction
type SaveUser struct {
	UserID *uuid.UUID `json:"user_id"`
}
