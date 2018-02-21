package infrastructure

import (
	"errors"
	"time"

	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	concrete_blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/infrastructure"
	validated "github.com/XMNBlockchain/core/packages/lives/blocks/validated/domain"
	hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/domain"
	concrete_hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/infrastructure"
	users "github.com/XMNBlockchain/core/packages/lives/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/lives/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

type blockBuilder struct {
	htBuilderFactory hashtrees.HashTreeBuilderFactory
	id               *uuid.UUID
	blk              blocks.SignedBlock
	ls               []users.Signature
	ts               *time.Time
}

func createBlockBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory) validated.BlockBuilder {
	out := blockBuilder{
		htBuilderFactory: htBuilderFactory,
		id:               nil,
		blk:              nil,
		ls:               nil,
		ts:               nil,
	}

	return &out
}

// Create initializes the BlockBuilder instance
func (build *blockBuilder) Create() validated.BlockBuilder {
	build.id = nil
	build.blk = nil
	build.ls = nil
	build.ts = nil
	return build
}

// WithID adds an ID to the BlockBuilder instance
func (build *blockBuilder) WithID(id *uuid.UUID) validated.BlockBuilder {
	build.id = id
	return build
}

// WithBlock adds a Block to the BlockBuilder instance
func (build *blockBuilder) WithBlock(blk blocks.SignedBlock) validated.BlockBuilder {
	build.blk = blk
	return build
}

// WithSignatures adds a leader signatures to the BlockBuilder instance
func (build *blockBuilder) WithSignatures(sigs []users.Signature) validated.BlockBuilder {
	build.ls = sigs
	return build
}

// CreatedOn adds a creation time to the BlockBuilder instance
func (build *blockBuilder) CreatedOn(ts time.Time) validated.BlockBuilder {
	build.ts = &ts
	return build
}

// Now builds a new Block instance
func (build *blockBuilder) Now() (validated.Block, error) {

	if build.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a validated block")
	}

	if build.blk == nil {
		return nil, errors.New("the block is mandatory in order to build a validated block")
	}

	if build.ls == nil {
		return nil, errors.New("the leader signatures are mandatory in order to build a block instance")
	}

	if len(build.ls) <= 0 {
		return nil, errors.New("the leader signatures cannot be empty in order to build a block instance")
	}

	if build.ts == nil {
		return nil, errors.New("the creation time is mandatory in order to build a validated block")
	}

	//add the block hashtree hash and the signature bytes as the first byte blocks:
	blks := [][]byte{
		build.blk.GetBlock().GetHashTree().GetHash().Get(),
		[]byte(build.blk.GetSignature().GetSig().String()),
	}

	//add each leader signature as a byte block:
	for _, oneSig := range build.ls {
		blks = append(blks, []byte(oneSig.GetSig().String()))
	}

	//build the hashtree with the byte blocks:
	ht, htErr := build.htBuilderFactory.Create().Create().WithBlocks(blks).Now()
	if htErr != nil {
		return nil, htErr
	}

	ls := []*concrete_users.Signature{}
	for _, oneSig := range build.ls {
		ls = append(ls, oneSig.(*concrete_users.Signature))
	}

	out := createBlock(build.id, ht.(*concrete_hashtrees.HashTree), build.blk.(*concrete_blocks.SignedBlock), ls, *build.ts)
	return out, nil
}
