package metadata

import (
	"errors"
	"time"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	uuid "github.com/satori/go.uuid"
)

type builder struct {
	id      *uuid.UUID
	crOn    *time.Time
	lstUpOn *time.Time
}

func createBuilder() metadata.Builder {
	out := builder{
		id:      nil,
		crOn:    nil,
		lstUpOn: nil,
	}

	return &out
}

// Create initializes the builder
func (build *builder) Create() metadata.Builder {
	build.id = nil
	build.crOn = nil
	build.lstUpOn = nil
	return build
}

// WithID adds an ID to the builder
func (build *builder) WithID(id *uuid.UUID) metadata.Builder {
	build.id = id
	return build
}

// CreatedOn adds a creation time to the builder
func (build *builder) CreatedOn(crOn time.Time) metadata.Builder {
	build.crOn = &crOn
	return build
}

// LastUpdatedOn adds a last updated on time to the builder
func (build *builder) LastUpdatedOn(lstUpOn time.Time) metadata.Builder {
	build.lstUpOn = &lstUpOn
	return build
}

// Now builds a new MetaData instance
func (build *builder) Now() (metadata.MetaData, error) {

	if build.lstUpOn == nil {
		build.lstUpOn = build.crOn
	}

	if build.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a MetaData instance")
	}

	if build.crOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a MetaData instance")
	}

	if build.lstUpOn.After(*build.crOn) {
		return nil, errors.New("the last updated on time cannot be before the creation time")
	}

	out := createMetaData(build.id, *build.crOn, *build.lstUpOn)
	return out, nil
}
