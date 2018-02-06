package custom

import (
	"encoding/json"
	"errors"

	custom "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/custom"
	uuid "github.com/satori/go.uuid"
)

type createBuilder struct {
	id  *uuid.UUID
	ins interface{}
}

func createCreateBuilder() custom.CreateBuilder {
	out := createBuilder{
		id:  nil,
		ins: nil,
	}

	return &out
}

// Create initializes the create builder
func (build *createBuilder) Create() custom.CreateBuilder {
	build.id = nil
	build.ins = nil
	return build
}

// WithID adds an ID the create builder
func (build *createBuilder) WithID(id *uuid.UUID) custom.CreateBuilder {
	build.id = id
	return build
}

// WithInstance adds an instance to the create builder
func (build *createBuilder) WithInstance(ins interface{}) custom.CreateBuilder {
	build.ins = ins
	return build
}

// Now builds a new Create instance
func (build *createBuilder) Now() (custom.Create, error) {
	if build.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a Create instance")
	}

	if build.ins == nil {
		return nil, errors.New("the instance is mandatory in order to build a Create instance")
	}

	js, jsErr := json.Marshal(build.ins)
	if jsErr != nil {
		return nil, jsErr
	}

	out := createCreate(build.id, js)
	return out, nil
}
