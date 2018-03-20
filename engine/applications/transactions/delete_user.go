package transactions

import uuid "github.com/satori/go.uuid"

// DeleteUser represents a delete user transaction
type DeleteUser struct {
	UserID *uuid.UUID `json:"user_id"`
}
