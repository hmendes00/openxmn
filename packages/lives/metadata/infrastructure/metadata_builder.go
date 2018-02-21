package infrastructure

import (
	"errors"
	"time"

	met "github.com/XMNBlockchain/core/packages/lives/metadata/domain"
	users "github.com/XMNBlockchain/core/packages/lives/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/lives/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

type metaDataBuilder struct {
	id   *uuid.UUID
	sig  users.Signature
	crOn *time.Time
}

func createMetaDataBuilder() met.MetaDataBuilder {
	out := metaDataBuilder{
		id:   nil,
		sig:  nil,
		crOn: nil,
	}
	return &out
}

// Create initializes a MetaDataBuilder instance
func (build *metaDataBuilder) Create() met.MetaDataBuilder {
	build.id = nil
	build.sig = nil
	build.crOn = nil
	return build
}

// WithID adds an ID to the MetaDataBuilder
func (build *metaDataBuilder) WithID(id *uuid.UUID) met.MetaDataBuilder {
	build.id = id
	return build
}

// WithSignature adds a signature to the MetaDataBuilder
func (build *metaDataBuilder) WithSignature(sig users.Signature) met.MetaDataBuilder {
	build.sig = sig
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

	if build.sig != nil {
		out := createMetaDataWithSignature(build.id, build.sig.(*concrete_users.Signature), *build.crOn)
		return out, nil
	}

	out := createMetaData(build.id, *build.crOn)
	return out, nil
}
