package infrastructure

import (
	"errors"
	"time"

	met "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	uuid "github.com/satori/go.uuid"
)

type metaDataBuilder struct {
	id   *uuid.UUID
	crOn *time.Time
}

func createMetaDataBuilder() met.MetaDataBuilder {
	out := metaDataBuilder{
		id:   nil,
		crOn: nil,
	}
	return &out
}

// Create initializes a MetaDataBuilder instance
func (build *metaDataBuilder) Create() met.MetaDataBuilder {
	build.id = nil
	build.crOn = nil
	return build
}

// WithID adds an ID to the MetaDataBuilder
func (build *metaDataBuilder) WithID(id *uuid.UUID) met.MetaDataBuilder {
	build.id = id
	return build
}

// CreatedOn adds a creation time
func (build *metaDataBuilder) CreatedOn(ts time.Time) met.MetaDataBuilder {
	build.crOn = &ts
	return build
}

// Now builds a new MetaData instance
func (build *metaDataBuilder) Now() (met.MetaData, error) {
	if build.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a MetaData instance")
	}

	if build.crOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a MetaData instance")
	}

	out := createMetaData(build.id, *build.crOn)
	return out, nil
}
