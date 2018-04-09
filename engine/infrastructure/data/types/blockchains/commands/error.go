package commands

import (
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	uuid "github.com/satori/go.uuid"
)

// Error represents a concrete error implementation
type Error struct {
	TrsID   *uuid.UUID `json:"transaction_id"`
	Code    int        `json:"code"`
	Message string     `json:"message"`
}

func createError(trsID *uuid.UUID, code int, message string) commands.Error {
	out := Error{
		TrsID:   trsID,
		Code:    code,
		Message: message,
	}

	return &out
}

// GetTransactionID returns the transactionID
func (err *Error) GetTransactionID() *uuid.UUID {
	return err.TrsID
}

// GetCode returns the code
func (err *Error) GetCode() int {
	return err.Code
}

// GetMessage returns the message
func (err *Error) GetMessage() string {
	return err.Message
}
