package infrastructure

import (
	"errors"

	chunks "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
	objs "github.com/XMNBlockchain/core/packages/lives/objects/domain"
)

type objectBuilder struct {
	metaData objs.MetaData
	chks     chunks.Chunks
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

// WithMetaData adds MetaData to the ObjectBuilder instance
func (build *objectBuilder) WithMetaData(met objs.MetaData) objs.ObjectBuilder {
	build.metaData = met
	return build
}

// WithChunks adds chunks to the ObjectBuilder
func (build *objectBuilder) WithChunks(chks chunks.Chunks) objs.ObjectBuilder {
	build.chks = chks
	return build
}

// Now creates an Object instance
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
