package validated

import (
	"errors"

	stored_block "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/blocks"
	stored_validated_block "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/blocks/validated"
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/users"
	concrete_stored_block "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/blockchains/blocks"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
	concrete_stored_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/blockchains/users"
)

type blockBuilder struct {
	metaData stored_files.File
	blk      stored_block.SignedBlock
	sigs     stored_users.Signatures
}

func createBlockBuilder() stored_validated_block.BlockBuilder {
	out := blockBuilder{
		metaData: nil,
		blk:      nil,
		sigs:     nil,
	}

	return &out
}

// Create initializes the BlockBuilder instance
func (build *blockBuilder) Create() stored_validated_block.BlockBuilder {
	build.metaData = nil
	build.blk = nil
	build.sigs = nil
	return build
}

// WithMetaData adds the metadata to the BlockBuilder instance
func (build *blockBuilder) WithMetaData(met stored_files.File) stored_validated_block.BlockBuilder {
	build.metaData = met
	return build
}

// WithBlock adds the block to the BlockBuilder instance
func (build *blockBuilder) WithBlock(blk stored_block.SignedBlock) stored_validated_block.BlockBuilder {
	build.blk = blk
	return build
}

// WithSignatures adds signatures to the BlockBuilder instance
func (build *blockBuilder) WithSignatures(sigs stored_users.Signatures) stored_validated_block.BlockBuilder {
	build.sigs = sigs
	return build
}

// Now builds a Block instance
func (build *blockBuilder) Now() (stored_validated_block.Block, error) {
	if build.metaData == nil {
		return nil, errors.New("the metadata is mandatory in order to build a Block instance")
	}

	if build.blk == nil {
		return nil, errors.New("the block is mandatory in order to build a Block instance")
	}

	if build.sigs == nil {
		return nil, errors.New("the signatures is mandatory in order to build a Block instance")
	}

	out := createBlock(build.metaData.(*concrete_stored_files.File), build.blk.(*concrete_stored_block.SignedBlock), build.sigs.(*concrete_stored_users.Signatures))
	return out, nil
}
