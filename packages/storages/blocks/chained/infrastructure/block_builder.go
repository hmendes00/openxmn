package infrastructure

import (
	"errors"

	stored_chained_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/chained/domain"
	stored_validated_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/validated/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

type blockBuilder struct {
	met stored_files.File
	ht  stored_files.File
	blk stored_validated_blocks.Block
}

func createBlockBuilder() stored_chained_blocks.BlockBuilder {
	out := blockBuilder{
		met: nil,
		ht:  nil,
		blk: nil,
	}

	return &out
}

// Create initializes the block builder
func (build *blockBuilder) Create() stored_chained_blocks.BlockBuilder {
	build.met = nil
	build.ht = nil
	build.blk = nil
	return build
}

// WithMetaData add the metadata file to the block builder
func (build *blockBuilder) WithMetaData(met stored_files.File) stored_chained_blocks.BlockBuilder {
	build.met = met
	return build
}

// WithHashTree add the hashtree file to the block builder
func (build *blockBuilder) WithHashTree(ht stored_files.File) stored_chained_blocks.BlockBuilder {
	build.ht = ht
	return build
}

// WithBlock add the stored block to the block builder
func (build *blockBuilder) WithBlock(blk stored_validated_blocks.Block) stored_chained_blocks.BlockBuilder {
	build.blk = blk
	return build
}

// Now builds a new stored chained block
func (build *blockBuilder) Now() (stored_chained_blocks.Block, error) {
	if build.met == nil {
		return nil, errors.New("the metadata file is mandatory in order to build a stored Block instance")
	}

	if build.ht == nil {
		return nil, errors.New("the hashtree file is mandatory in order to build a stored Block instance")
	}

	if build.blk == nil {
		return nil, errors.New("the stored validated block is mandatory in order to build a stored Block instance")
	}

	out := createBlock(build.met, build.ht, build.blk)
	return out, nil
}
