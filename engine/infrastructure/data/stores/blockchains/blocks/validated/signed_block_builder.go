package validated

import (
	"errors"

	stored_validated_block "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/blocks/validated"
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/users"
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	concrete_stored_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/blockchains/users"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
)

type signedBlockBuilder struct {
	met stored_files.File
	blk stored_validated_block.Block
	sig stored_users.Signature
}

func createSignedBlockBuilder() stored_validated_block.SignedBlockBuilder {
	out := signedBlockBuilder{
		met: nil,
		blk: nil,
		sig: nil,
	}

	return &out
}

// Create intializes the SignedBlockBuilder instance
func (build *signedBlockBuilder) Create() stored_validated_block.SignedBlockBuilder {
	build.met = nil
	build.blk = nil
	build.sig = nil
	return build
}

// WithMetaData adds metadata to the SignedBlockBuilder instance
func (build *signedBlockBuilder) WithMetaData(met stored_files.File) stored_validated_block.SignedBlockBuilder {
	build.met = met
	return build
}

// WithBlock adds block to the SignedBlockBuilder instance
func (build *signedBlockBuilder) WithBlock(blk stored_validated_block.Block) stored_validated_block.SignedBlockBuilder {
	build.blk = blk
	return build
}

// WithSignature adds a signature to the SignedBlockBuilder instance
func (build *signedBlockBuilder) WithSignature(sig stored_users.Signature) stored_validated_block.SignedBlockBuilder {
	build.sig = sig
	return build
}

// Now builds a new SignedBlock instance
func (build *signedBlockBuilder) Now() (stored_validated_block.SignedBlock, error) {
	if build.met == nil {
		return nil, errors.New("the metadata is mandatory in order to build a SignedBlock instance")
	}

	if build.blk == nil {
		return nil, errors.New("the block is mandatory in order to build a SignedBlock instance")
	}

	if build.sig == nil {
		return nil, errors.New("the signature is mandatory in order to build a SignedBlock instance")
	}

	out := createSignedBlock(build.met.(*concrete_stored_files.File), build.blk.(*Block), build.sig.(*concrete_stored_users.Signature))
	return out, nil
}
