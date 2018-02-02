package users

import uuid "github.com/satori/go.uuid"

// DeleteBuilder represents the builder of a delete user transaction
type DeleteBuilder interface {
	Create() DeleteBuilder
	WithID(id *uuid.UUID) DeleteBuilder
	Now() (Delete, error)
}
