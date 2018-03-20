package metadata

import (
	"errors"
	"time"

	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
	met "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	concrete_hashtrees "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/hashtrees"
	uuid "github.com/satori/go.uuid"
)

type metaDataBuilder struct {
	id   *uuid.UUID
	ht   hashtrees.HashTree
	crOn *time.Time
}

func createMetaDataBuilder() met.MetaDataBuilder {
	out := metaDataBuilder{
		id:   nil,
		ht:   nil,
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

// WithHashTree adds an hashtree to the MetaDataBuilder
func (build *metaDataBuilder) WithHashTree(ht hashtrees.HashTree) met.MetaDataBuilder {
	build.ht = ht
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

	if build.ht == nil {
		return nil, errors.New("the HashTree is mandatory in order to build a MetaData instance")
	}

	if build.crOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a MetaData instance")
	}

	out := createMetaData(build.id, build.ht.(*concrete_hashtrees.HashTree), *build.crOn)
	return out, nil
}
