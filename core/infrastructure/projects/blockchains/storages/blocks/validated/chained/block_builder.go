package chained

import (
	"errors"

	stored_validated_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks/validated"
	stored_chained_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks/validated/chained"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	concrete_stored_validated_blocks "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/blocks/validated"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/files"
)

type blockBuilder struct {
	met stored_files.File
	blk stored_validated_blocks.Block
}

func createBlockBuilder() stored_chained_blocks.BlockBuilder {
	out := blockBuilder{
		met: nil,
		blk: nil,
	}

	return &out
}

// Create initializes the block builder
func (build *blockBuilder) Create() stored_chained_blocks.BlockBuilder {
	build.met = nil
	build.blk = nil
	return build
}

// WithMetaData add the metadata file to the block builder
func (build *blockBuilder) WithMetaData(met stored_files.File) stored_chained_blocks.BlockBuilder {
	build.met = met
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

	if build.blk == nil {
		return nil, errors.New("the stored validated block is mandatory in order to build a stored Block instance")
	}

	out := createBlock(build.met.(*concrete_stored_files.File), build.blk.(*concrete_stored_validated_blocks.Block))
	return out, nil
}
