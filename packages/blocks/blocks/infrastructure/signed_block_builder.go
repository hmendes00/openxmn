package infrastructure

import (
	"errors"

	blocks "github.com/XMNBlockchain/core/packages/blocks/blocks/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/users/infrastructure"
)

type signedBlockBuilder struct {
	blk blocks.Block
	sig users.Signature
}

func createSignedBlockBuilder() blocks.SignedBlockBuilder {
	out := signedBlockBuilder{
		blk: nil,
		sig: nil,
	}

	return &out
}

// Create initializes the SignedBlockBuilder instance
func (build *signedBlockBuilder) Create() blocks.SignedBlockBuilder {
	build.blk = nil
	build.sig = nil
	return build
}

// WithBlock adds a Block to the SignedBlockBuilder instance
func (build *signedBlockBuilder) WithBlock(blk blocks.Block) blocks.SignedBlockBuilder {
	build.blk = blk
	return build
}

// WithSignature adds a user Signature to the SignedBlockBuilder instance
func (build *signedBlockBuilder) WithSignature(sig users.Signature) blocks.SignedBlockBuilder {
	build.sig = sig
	return build
}

// Now builds a SignedBlock instance
func (build *signedBlockBuilder) Now() (blocks.SignedBlock, error) {

	if build.blk == nil {
		return nil, errors.New("the block is mandatory in order to build a SignedBlock instance")
	}

	if build.sig == nil {
		return nil, errors.New("the user signature is mandatory in order to build a SignedBlock instance")
	}

	out := createSignedBlock(build.blk.(*Block), build.sig.(*concrete_users.Signature))
	return out, nil
}
