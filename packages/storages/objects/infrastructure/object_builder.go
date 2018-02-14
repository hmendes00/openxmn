package infrastructure

import (
	"errors"

	stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	objs "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

type objectBuilder struct {
	metaData stored_files.File
	chks     stored_chunks.Chunks
}

func createObjectBuilder() objs.ObjectBuilder {
	out := objectBuilder{
		metaData: nil,
		chks:     nil,
	}

	return &out
}

// Create initializes the ObjectBuilder
func (build *objectBuilder) Create() objs.ObjectBuilder {
	build.metaData = nil
	build.chks = nil
	return build
}

// WithMetaData adds a MetaData to the ObjectBuilder instance
func (build *objectBuilder) WithMetaData(met stored_files.File) objs.ObjectBuilder {
	build.metaData = met
	return build
}

// WithChunks adds a stored chunks to the ObjectBuilder
func (build *objectBuilder) WithChunks(chks stored_chunks.Chunks) objs.ObjectBuilder {
	build.chks = chks
	return build
}

// Now builds a new Object instance
func (build *objectBuilder) Now() (objs.Object, error) {

	if build.metaData == nil {
		return nil, errors.New("the MetaData is mandatory in order to build an Object instance")
	}

	if build.chks != nil {
		out := createObjectWithChunks(build.metaData, build.chks)
		return out, nil
	}

	out := createObject(build.metaData)
	return out, nil
}
