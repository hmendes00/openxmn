package infrastructure

import (
	"errors"

	stored_block "github.com/XMNBlockchain/core/packages/storages/blocks/blocks/domain"
	stored_validated_block "github.com/XMNBlockchain/core/packages/storages/blocks/validated/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

type blockBuilder struct {
	metaData stored_files.File
	ht       stored_files.File
	blk      stored_block.SignedBlock
	sigs     []stored_files.File
}

func createBlockBuilder() stored_validated_block.BlockBuilder {
	out := blockBuilder{
		metaData: nil,
		ht:       nil,
		blk:      nil,
		sigs:     nil,
	}

	return &out
}

// Create initializes the BlockBuilder instance
func (build *blockBuilder) Create() stored_validated_block.BlockBuilder {
	build.metaData = nil
	build.ht = nil
	build.blk = nil
	build.sigs = nil
	return build
}

// WithMetaData adds the metadata to the BlockBuilder instance
func (build *blockBuilder) WithMetaData(met stored_files.File) stored_validated_block.BlockBuilder {
	build.metaData = met
	return build
}

// WithHashTree adds the hashtree to the BlockBuilder instance
func (build *blockBuilder) WithHashTree(ht stored_files.File) stored_validated_block.BlockBuilder {
	build.ht = ht
	return build
}

// WithBlock adds the block to the BlockBuilder instance
func (build *blockBuilder) WithBlock(blk stored_block.SignedBlock) stored_validated_block.BlockBuilder {
	build.blk = blk
	return build
}

// WithSignatures adds signatures to the BlockBuilder instance
func (build *blockBuilder) WithSignatures(sigs []stored_files.File) stored_validated_block.BlockBuilder {
	build.sigs = sigs
	return build
}

// Now builds a Block instance
func (build *blockBuilder) Now() (stored_validated_block.Block, error) {
	if build.metaData == nil {
		return nil, errors.New("the metadata is mandatory in order to build a Block instance")
	}

	if build.ht == nil {
		return nil, errors.New("the hashtree is mandatory in order to build a Block instance")
	}

	if build.blk == nil {
		return nil, errors.New("the block is mandatory in order to build a Block instance")
	}

	if build.sigs == nil {
		return nil, errors.New("the signatures is mandatory in order to build a Block instance")
	}

	out := createBlock(build.metaData, build.ht, build.blk, build.sigs)
	return out, nil
}
