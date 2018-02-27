package infrastructure

import (
	"errors"
	"time"

	blocks "github.com/XMNBlockchain/core/packages/blockchains/blocks/blocks/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/blockchains/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

type signedBlockBuilder struct {
	id        *uuid.UUID
	blk       blocks.Block
	sig       users.Signature
	createdOn *time.Time
}

func createSignedBlockBuilder() blocks.SignedBlockBuilder {
	out := signedBlockBuilder{
		id:        nil,
		blk:       nil,
		sig:       nil,
		createdOn: nil,
	}

	return &out
}

// Create initializes the SignedBlockBuilder instance
func (build *signedBlockBuilder) Create() blocks.SignedBlockBuilder {
	build.id = nil
	build.blk = nil
	build.sig = nil
	build.createdOn = nil
	return build
}

// WithID adds the ID to the SignedBlockBuilder instance
func (build *signedBlockBuilder) WithID(id *uuid.UUID) blocks.SignedBlockBuilder {
	build.id = id
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

// CreatedOn adds the creation time to the SignedBlockBuilder instance
func (build *signedBlockBuilder) CreatedOn(createdOn time.Time) blocks.SignedBlockBuilder {
	build.createdOn = &createdOn
	return build
}

// Now builds a SignedBlock instance
func (build *signedBlockBuilder) Now() (blocks.SignedBlock, error) {

	if build.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a SignedBlock instance")
	}

	if build.blk == nil {
		return nil, errors.New("the block is mandatory in order to build a SignedBlock instance")
	}

	if build.sig == nil {
		return nil, errors.New("the user signature is mandatory in order to build a SignedBlock instance")
	}

	if build.createdOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a SignedBlock instance")
	}

	out := createSignedBlock(build.id, build.blk.(*Block), build.sig.(*concrete_users.Signature), *build.createdOn)
	return out, nil
}
