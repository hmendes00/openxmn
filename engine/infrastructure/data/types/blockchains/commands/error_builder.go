package commands

import (
	"errors"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	uuid "github.com/satori/go.uuid"
)

type errorBuilder struct {
	trsID   *uuid.UUID
	code    int
	message string
}

func createErrorBuilder() commands.ErrorBuilder {
	out := errorBuilder{
		trsID:   nil,
		code:    0,
		message: "",
	}

	return &out
}

// Create initializes the ErrorBuilder instance
func (build *errorBuilder) Create() commands.ErrorBuilder {
	build.trsID = nil
	build.code = 0
	build.message = ""
	return build
}

// WithTransactionID adds a transactionID to the ErrorBuilder instance
func (build *errorBuilder) WithTransactionID(transID *uuid.UUID) commands.ErrorBuilder {
	build.trsID = transID
	return build
}

// WithMessage adds a message to the ErrorBuilder instance
func (build *errorBuilder) WithMessage(msg string) commands.ErrorBuilder {
	build.message = msg
	return build
}

// WithCode adds a code to the ErrorBuilder instance
func (build *errorBuilder) WithCode(code int) commands.ErrorBuilder {
	build.code = code
	return build
}

// Now builds a new Error instance
func (build *errorBuilder) Now() (commands.Error, error) {
	if build.trsID == nil {
		return nil, errors.New("the transactionID is mandatory in order to build an Error instance")
	}

	if build.message == "" {
		return nil, errors.New("the message is mandatory in order to build an Error instance")
	}

	if build.code == 0 {
		return nil, errors.New("the code is mandatory in order to build an Error instance")
	}

	out := createError(build.trsID, build.code, build.message)
	return out, nil
}
