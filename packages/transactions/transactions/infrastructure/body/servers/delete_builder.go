package servers

import (
	"errors"

	servers "github.com/XMNBlockchain/core/packages/transactions/transactions/domain/body/servers"
	uuid "github.com/satori/go.uuid"
)

type deleteBuilder struct {
	id *uuid.UUID
}

func createDeleteBuilder() servers.DeleteBuilder {
	out := deleteBuilder{
		id: nil,
	}

	return &out
}

// Create initializes the DeleteBuilder
func (build *deleteBuilder) Create() servers.DeleteBuilder {
	build.id = nil
	return build
}

// WithID adds an ID to the DeleteBuilder
func (build *deleteBuilder) WithID(id *uuid.UUID) servers.DeleteBuilder {
	build.id = id
	return build
}

// Now builds a new Delete instance
func (build *deleteBuilder) Now() (servers.Delete, error) {

	if build.id == nil {
		return nil, errors.New("the id is mandatory in order to build a Delete instance")
	}

	out := createDelete(build.id)
	return out, nil
}
