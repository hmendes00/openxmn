package infrastructure

import (
	"errors"

	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
)

type hashTreeBuilder struct {
	blocks [][]byte
	js     []byte
}

func createHashTreeBuilder() hashtrees.HashTreeBuilder {
	out := hashTreeBuilder{}
	return &out
}

// Create initializes the HashTreeBuilder instance
func (build *hashTreeBuilder) Create() hashtrees.HashTreeBuilder {
	build.blocks = nil
	build.js = nil
	return build
}

// WithBlocks adds blocks to the HashTreeBuilder
func (build *hashTreeBuilder) WithBlocks(blocks [][]byte) hashtrees.HashTreeBuilder {
	build.blocks = blocks
	return build
}

// WithJSON adds JSON to the HashTreeBuilder
func (build *hashTreeBuilder) WithJSON(js []byte) hashtrees.HashTreeBuilder {
	build.js = js
	return build
}

// Now builds an HashTree instance
func (build *hashTreeBuilder) Now() (hashtrees.HashTree, error) {

	if build.blocks != nil && build.js != nil {
		return nil, errors.New("the blocks or the js must be set.  both set")
	}

	if build.blocks != nil {
		out, outErr := createHashTreeFromBlocks(build.blocks)
		return out, outErr
	}

	if build.js != nil {
		out, outErr := createHashTreeFromJSON(build.js)
		return out, outErr
	}

	return nil, errors.New("the blocks or js must be set.  both nil")

}
