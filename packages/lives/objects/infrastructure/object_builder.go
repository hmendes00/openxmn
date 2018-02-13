package infrastructure

import (
	"errors"
	"time"

	chunks "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
	objs "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	uuid "github.com/satori/go.uuid"
)

type objectBuilder struct {
	id   *uuid.UUID
	crOn *time.Time
	sig  users.Signature
	chks chunks.Chunks
}

func createObjectBuilder() objs.ObjectBuilder {
	out := objectBuilder{
		id:   nil,
		crOn: nil,
		sig:  nil,
		chks: nil,
	}

	return &out
}

// Create initializes the ObjectBuilder
func (build *objectBuilder) Create() objs.ObjectBuilder {
	build.id = nil
	build.crOn = nil
	build.sig = nil
	build.chks = nil
	return build
}

// WithID adds an ID to the ObjectBuilder
func (build *objectBuilder) WithID(id *uuid.UUID) objs.ObjectBuilder {
	build.id = id
	return build
}

// WithSignature adds a signature to the ObjectBuilder
func (build *objectBuilder) WithSignature(sig users.Signature) objs.ObjectBuilder {
	build.sig = sig
	return build
}

// WithChunks adds chunks to the ObjectBuilder
func (build *objectBuilder) WithChunks(chks chunks.Chunks) objs.ObjectBuilder {
	build.chks = chks
	return build
}

// CreatedOn adds creation time to the ObjectBuilder
func (build *objectBuilder) CreatedOn(ts time.Time) objs.ObjectBuilder {
	build.crOn = &ts
	return build
}

// Now creates an Object instance
func (build *objectBuilder) Now() (objs.Object, error) {
	if build.id == nil {
		return nil, errors.New("the ID is mandatory in order to build an Object instance")
	}

	if build.crOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build an Object instance")
	}

	if build.chks != nil && build.sig != nil {
		out := createObjectWithChunksWithSignature(build.id, *build.crOn, build.chks, build.sig)
		return out, nil
	}

	if build.chks != nil {
		out := createObjectWithChunks(build.id, *build.crOn, build.chks)
		return out, nil
	}

	if build.sig != nil {
		out := createObjectWithSignature(build.id, *build.crOn, build.sig)
		return out, nil
	}

	out := createObject(build.id, *build.crOn)
	return out, nil
}
