package infrastructure

import (
	"errors"
	"time"

	chained "github.com/XMNBlockchain/core/packages/blockchains/blocks/chained/domain"
	uuid "github.com/satori/go.uuid"
)

type metaDataBuilder struct {
	id        *uuid.UUID
	index     int
	prevIndex int
	crOn      *time.Time
}

func createMetaDataBuilder() chained.MetaDataBuilder {
	out := metaDataBuilder{
		id:        nil,
		index:     0,
		prevIndex: -1,
		crOn:      nil,
	}

	return &out
}

// Create initializes the MetaData builder
func (build *metaDataBuilder) Create() chained.MetaDataBuilder {
	build.id = nil
	build.index = 0
	build.prevIndex = -1
	build.crOn = nil
	return build
}

// WithID adds an ID to the metadata builder
func (build *metaDataBuilder) WithID(id *uuid.UUID) chained.MetaDataBuilder {
	build.id = id
	return build
}

// WithIndex adds an index to the metadata builder
func (build *metaDataBuilder) WithIndex(index int) chained.MetaDataBuilder {
	build.index = index
	return build
}

// WithPreviousIndex adds a previous index to the metadata builder
func (build *metaDataBuilder) WithPreviousIndex(prevIndex int) chained.MetaDataBuilder {
	build.prevIndex = prevIndex
	return build
}

// CreatedOn adds a creation time
func (build *metaDataBuilder) CreatedOn(crOn time.Time) chained.MetaDataBuilder {
	build.crOn = &crOn
	return build
}

// Now builds a new MetaData instance
func (build *metaDataBuilder) Now() (chained.MetaData, error) {
	if build.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a MetaData instance")
	}

	if build.index == 0 {
		return nil, errors.New("the index is mandatory in order to build a MetaData instance")
	}

	if build.prevIndex == -1 {
		return nil, errors.New("the previous index is mandatory in order to build a MetaData instance")
	}

	if build.crOn == nil {
		return nil, errors.New("the createdOn time is mandatory in order to build a MetaData instance")
	}

	out := createMetaData(build.id, build.index, build.prevIndex, *build.crOn)
	return out, nil
}
