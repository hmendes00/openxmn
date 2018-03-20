package blocks

import (
	"errors"

	stored_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/blocks"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/files"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/users"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/storages/files"
	concrete_stored_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/storages/users"
)

type signedBlockBuilder struct {
	metaData stored_files.File
	sig      stored_users.Signature
	blk      stored_blocks.Block
}

func createSignedBlockBuilder() stored_blocks.SignedBlockBuilder {
	out := signedBlockBuilder{
		metaData: nil,
		sig:      nil,
		blk:      nil,
	}

	return &out
}

// Create initializes the SignedBlockBuilder instance
func (build *signedBlockBuilder) Create() stored_blocks.SignedBlockBuilder {
	build.metaData = nil
	build.sig = nil
	build.blk = nil
	return build
}

// WithMetaData adds metadata to the SignedBlockBuilder instance
func (build *signedBlockBuilder) WithMetaData(met stored_files.File) stored_blocks.SignedBlockBuilder {
	build.metaData = met
	return build
}

// WithSignature adds a signature to the SignedBlockBuilder instance
func (build *signedBlockBuilder) WithSignature(sig stored_users.Signature) stored_blocks.SignedBlockBuilder {
	build.sig = sig
	return build
}

// WithBlock adds a stored block to the SignedBlockBuilder instance
func (build *signedBlockBuilder) WithBlock(blk stored_blocks.Block) stored_blocks.SignedBlockBuilder {
	build.blk = blk
	return build
}

// Now builds a new SignedBlock instance
func (build *signedBlockBuilder) Now() (stored_blocks.SignedBlock, error) {
	if build.metaData == nil {
		return nil, errors.New("the metadata is mandatory in order to build a SignedBlock instance")
	}

	if build.sig == nil {
		return nil, errors.New("the signature is mandatory in order to build a SignedBlock instance")
	}

	if build.blk == nil {
		return nil, errors.New("the block is mandatory in order to build a SignedBlock instance")
	}

	out := createSignedBlock(build.metaData.(*concrete_stored_files.File), build.sig.(*concrete_stored_users.Signature), build.blk.(*Block))
	return out, nil
}
