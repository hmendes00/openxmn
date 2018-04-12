package wealth

import uuid "github.com/satori/go.uuid"

// DeleteUser represents a delete user transaction
type DeleteUser struct {
	UserID *uuid.UUID `json:"user_id"`
}

// CreateDeleteUser creates a new DeleteUser instance
func CreateDeleteUser(userID *uuid.UUID) *DeleteUser {
	out := DeleteUser{
		UserID: userID,
	}

	return &out
}

// GetUserID returns the userID
func (del *DeleteUser) GetUserID() *uuid.UUID {
	return del.UserID
}
