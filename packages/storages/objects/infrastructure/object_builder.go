package infrastructure

import (
	"errors"
	"time"

	stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
	uuid "github.com/satori/go.uuid"
)

type objectBuilder struct {
	ht   stored_files.File
	id   *uuid.UUID
	crOn *time.Time
	sig  stored_files.File
	chks stored_chunks.Chunks
}

func createObjectBuilder() objects.ObjectBuilder {
	out := objectBuilder{
		ht:   nil,
		id:   nil,
		crOn: nil,
		sig:  nil,
		chks: nil,
	}

	return &out
}

// Create initializes the ObjectBuilder
func (build *objectBuilder) Create() objects.ObjectBuilder {
	build.ht = nil
	build.id = nil
	build.crOn = nil
	build.sig = nil
	build.chks = nil
	return build
}

// WithHashTree adds an HashTree file to the ObjectBuilder
func (build *objectBuilder) WithHashTree(ht stored_files.File) objects.ObjectBuilder {
	build.ht = ht
	return build
}

// WithID adds an ID to the ObjectBuilder
func (build *objectBuilder) WithID(id *uuid.UUID) objects.ObjectBuilder {
	build.id = id
	return build
}

// WithSignature adds a signature file to the ObjectBuilder
func (build *objectBuilder) WithSignature(sig stored_files.File) objects.ObjectBuilder {
	build.sig = sig
	return build
}

// WithChunks adds a stored chunks to the ObjectBuilder
func (build *objectBuilder) WithChunks(chks stored_chunks.Chunks) objects.ObjectBuilder {
	build.chks = chks
	return build
}

// CreatedOn adds a creation time to the ObjectBuilder
func (build *objectBuilder) CreatedOn(ts time.Time) objects.ObjectBuilder {
	build.crOn = &ts
	return build
}

// Now builds a new Object instance
func (build *objectBuilder) Now() (objects.Object, error) {
	if build.ht == nil {
		return nil, errors.New("the hashtree is mandatory in order to build an Object instance")
	}

	if build.id == nil {
		return nil, errors.New("the ID is mandatory in order to build an Object instance")
	}

	if build.crOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build an Object instance")
	}

	if build.sig != nil && build.chks != nil {
		out := createObjectWithSignatureWithChunks(build.ht, build.id, *build.crOn, build.sig, build.chks)
		return out, nil
	}

	if build.sig != nil {
		out := createObjectWithSignature(build.ht, build.id, *build.crOn, build.sig)
		return out, nil
	}

	if build.chks != nil {
		out := createObjectWithChunks(build.ht, build.id, *build.crOn, build.chks)
		return out, nil
	}

	out := createObject(build.ht, build.id, *build.crOn)
	return out, nil
}
