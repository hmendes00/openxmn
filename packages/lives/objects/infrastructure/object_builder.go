package infrastructure

import (
	"errors"
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	chunks "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	uuid "github.com/satori/go.uuid"
)

type objectBuilder struct {
	id   *uuid.UUID
	ht   hashtrees.HashTree
	crOn *time.Time
	sig  users.Signature
	chks chunks.Chunks
}

func createObjectBuilder() objects.ObjectBuilder {
	out := objectBuilder{
		id:   nil,
		ht:   nil,
		crOn: nil,
		sig:  nil,
		chks: nil,
	}

	return &out
}

// Create initializes the ObjectBuilder
func (build *objectBuilder) Create() objects.ObjectBuilder {
	build.id = nil
	build.ht = nil
	build.crOn = nil
	build.sig = nil
	build.chks = nil
	return build
}

// WithID adds an ID to the ObjectBuilder
func (build *objectBuilder) WithID(id *uuid.UUID) objects.ObjectBuilder {
	build.id = id
	return build
}

// WithSignature adds a signature to the ObjectBuilder
func (build *objectBuilder) WithSignature(sig users.Signature) objects.ObjectBuilder {
	build.sig = sig
	return build
}

// WithChunks adds chunks to the ObjectBuilder
func (build *objectBuilder) WithChunks(chks chunks.Chunks) objects.ObjectBuilder {
	build.chks = chks
	return build
}

// CreatedOn adds creation time to the ObjectBuilder
func (build *objectBuilder) CreatedOn(ts time.Time) objects.ObjectBuilder {
	build.crOn = &ts
	return build
}

// Now creates an Object instance
func (build *objectBuilder) Now() (objects.Object, error) {
	if build.id == nil {
		return nil, errors.New("the ID is mandatory in order to build an Object instance")
	}

	if build.chks == nil {
		return nil, errors.New("the chunks are mandatory in order to build an Object instance")
	}

	if build.crOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build an Object instance")
	}

	if build.sig != nil {
		out := createObjectWithSignature(build.id, build.chks, build.sig, *build.crOn)
		return out, nil
	}

	out := createObject(build.id, build.chks, *build.crOn)
	return out, nil
}
