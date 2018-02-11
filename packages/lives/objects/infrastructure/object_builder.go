package infrastructure

import (
	"errors"
	"strconv"
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	chunks "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	uuid "github.com/satori/go.uuid"
)

type objectBuilder struct {
	htBuilderFactory hashtrees.HashTreeBuilderFactory
	id               *uuid.UUID
	path             string
	ht               hashtrees.HashTree
	crOn             *time.Time
	sig              users.Signature
	chks             chunks.Chunks
}

func createObjectBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory) objects.ObjectBuilder {
	out := objectBuilder{
		htBuilderFactory: htBuilderFactory,
		id:               nil,
		path:             "",
		ht:               nil,
		crOn:             nil,
		sig:              nil,
		chks:             nil,
	}

	return &out
}

// Create initializes the ObjectBuilder
func (build *objectBuilder) Create() objects.ObjectBuilder {
	build.id = nil
	build.path = ""
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

// WithPath adds a path to the ObjectBuilder
func (build *objectBuilder) WithPath(path string) objects.ObjectBuilder {
	build.path = path
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

	if build.path == "" {
		return nil, errors.New("the path is mandatory in order to build an Object instance")
	}

	if build.crOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build an Object instance")
	}

	crOnAsString := strconv.Itoa(int(build.crOn.Unix()))

	htData := [][]byte{
		build.id.Bytes(),
		[]byte(build.path),
		[]byte(crOnAsString),
	}

	if build.chks != nil {
		htData = append(htData, build.chks.GetHashTree().GetHash().Get())
	}

	if build.sig != nil {
		htData = append(htData, []byte(build.sig.GetSig().String()))
	}

	//build the hash tree:
	ht, htErr := build.htBuilderFactory.Create().Create().WithBlocks(htData).Now()
	if htErr != nil {
		return nil, htErr
	}

	if build.chks != nil && build.sig != nil {
		out := createObjectWithChunksWithSignature(build.id, build.path, ht, *build.crOn, build.chks, build.sig)
		return out, nil
	}

	if build.chks != nil {
		out := createObjectWithSignature(build.id, build.path, ht, *build.crOn, build.sig)
		return out, nil
	}

	if build.sig != nil {
		out := createObjectWithSignature(build.id, build.path, ht, *build.crOn, build.sig)
		return out, nil
	}

	out := createObject(build.id, build.path, ht, *build.crOn)
	return out, nil
}
