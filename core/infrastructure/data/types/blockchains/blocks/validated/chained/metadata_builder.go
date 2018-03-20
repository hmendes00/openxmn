package chained

import (
	"errors"
	"time"

	chained "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/blocks/validated/chained"
	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/hashtrees"
	concrete_hashtrees "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/hashtrees"
	uuid "github.com/satori/go.uuid"
)

type metaDataBuilder struct {
	id     *uuid.UUID
	ht     hashtrees.HashTree
	prevID *uuid.UUID
	crOn   *time.Time
}

func createMetaDataBuilder() chained.MetaDataBuilder {
	out := metaDataBuilder{
		id:     nil,
		ht:     nil,
		prevID: nil,
		crOn:   nil,
	}

	return &out
}

// Create initializes the MetaData builder
func (build *metaDataBuilder) Create() chained.MetaDataBuilder {
	build.id = nil
	build.ht = nil
	build.prevID = nil
	build.crOn = nil
	return build
}

// WithID adds an ID to the metadata builder
func (build *metaDataBuilder) WithID(id *uuid.UUID) chained.MetaDataBuilder {
	build.id = id
	return build
}

// WithHashTree adds an HashTree to the metadata builder
func (build *metaDataBuilder) WithHashTree(ht hashtrees.HashTree) chained.MetaDataBuilder {
	build.ht = ht
	return build
}

// WithPreviousID adds a previous ID to the metadata builder
func (build *metaDataBuilder) WithPreviousID(prevID *uuid.UUID) chained.MetaDataBuilder {
	build.prevID = prevID
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

	if build.ht == nil {
		return nil, errors.New("the HashTree is mandatory in order to build a MetaData instance")
	}

	if build.prevID == nil {
		return nil, errors.New("the previous ID is mandatory in order to build a MetaData instance")
	}

	if build.crOn == nil {
		return nil, errors.New("the createdOn time is mandatory in order to build a MetaData instance")
	}

	out := createMetaData(build.id, build.ht.(*concrete_hashtrees.HashTree), build.prevID, *build.crOn)
	return out, nil
}
