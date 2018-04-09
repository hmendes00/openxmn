package commands

import uuid "github.com/satori/go.uuid"

// Error represents an error command
type Error interface {
	GetTransactionID() *uuid.UUID
	GetCode() int
	GetMessage() string
}
