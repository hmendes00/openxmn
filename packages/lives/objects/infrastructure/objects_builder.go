package infrastructure

import (
	"errors"

	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	objs "github.com/XMNBlockchain/core/packages/lives/objects/domain"
)

type objectsBuilder struct {
	htBuilderFactory hashtrees.HashTreeBuilderFactory
	objs             []objs.Object
}

func createObjectsBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory) objs.ObjectsBuilder {
	out := objectsBuilder{
		htBuilderFactory: htBuilderFactory,
		objs:             nil,
	}

	return &out
}

// Create initializes an ObjectsBuilder instance
func (build *objectsBuilder) Create() objs.ObjectsBuilder {
	build.objs = nil
	return build
}

// WithObjects adds []Object to the ObjectsBuilder instance
func (build *objectsBuilder) WithObjects(objs []objs.Object) objs.ObjectsBuilder {
	build.objs = objs
	return build
}

// Now builds a new Objects instance
func (build *objectsBuilder) Now() (objs.Objects, error) {

	if build.objs == nil {
		return nil, errors.New("the []Object are mandatory in order to build an Objects instance")
	}

	if len(build.objs) <= 0 {
		return nil, errors.New("the []Object must contain at least 1 instance, none provided")
	}

	//create the blocks:
	blocks := [][]byte{}
	for _, oneObj := range build.objs {
		idAsBytes := oneObj.GetMetaData().GetID().Bytes()
		blocks = append(blocks, idAsBytes)
	}

	ht, htErr := build.htBuilderFactory.Create().Create().WithBlocks(blocks).Now()
	if htErr != nil {
		return nil, htErr
	}

	out := createObjects(ht, build.objs)
	return out, nil
}
