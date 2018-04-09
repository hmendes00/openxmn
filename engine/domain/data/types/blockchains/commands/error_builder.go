package commands

import uuid "github.com/satori/go.uuid"

// ErrorBuilder represents an error builder
type ErrorBuilder interface {
	Create() ErrorBuilder
	WithTransactionID(transID *uuid.UUID) ErrorBuilder
	WithMessage(msg string) ErrorBuilder
	WithCode(code int) ErrorBuilder
	Now() (Error, error)
}
