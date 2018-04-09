package metadata

import (
	"errors"
	"time"

	met "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
	concrete_hashtrees "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/hashtrees"
	uuid "github.com/satori/go.uuid"
)

type metaDataBuilder struct {
	id   *uuid.UUID
	ht   hashtrees.HashTree
	crOn *time.Time
}

func createBuilder() met.Builder {
	out := metaDataBuilder{
		id:   nil,
		ht:   nil,
		crOn: nil,
	}
	return &out
}

// Create initializes a Builder instance
func (build *metaDataBuilder) Create() met.Builder {
	build.id = nil
	build.crOn = nil
	return build
}

// WithID adds an ID to the Builder
func (build *metaDataBuilder) WithID(id *uuid.UUID) met.Builder {
	build.id = id
	return build
}

// WithHashTree adds an hashtree to the Builder
func (build *metaDataBuilder) WithHashTree(ht hashtrees.HashTree) met.Builder {
	build.ht = ht
	return build
}

// CreatedOn adds a creation time
func (build *metaDataBuilder) CreatedOn(ts time.Time) met.Builder {
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
