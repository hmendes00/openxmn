package hashtrees

import (
	"errors"

	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
)

type hashTreeBuilder struct {
	blocks [][]byte
}

func createHashTreeBuilder() hashtrees.HashTreeBuilder {
	out := hashTreeBuilder{}
	return &out
}

// Create initializes the HashTreeBuilder instance
func (build *hashTreeBuilder) Create() hashtrees.HashTreeBuilder {
	build.blocks = nil
	return build
}

// WithBlocks adds blocks to the HashTreeBuilder
func (build *hashTreeBuilder) WithBlocks(blocks [][]byte) hashtrees.HashTreeBuilder {
	build.blocks = blocks
	return build
}

// Now builds an HashTree instance
func (build *hashTreeBuilder) Now() (hashtrees.HashTree, error) {

	if build.blocks == nil {
		return nil, errors.New("the blocks are mandatory in order to build an HashTree instance")
	}

	out, outErr := createHashTreeFromBlocks(build.blocks)
	return out, outErr

}
