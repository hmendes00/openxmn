package infrastructure

import (
	"errors"

	blocks "github.com/XMNBlockchain/core/packages/blocks/blocks/domain"
	concrete_blocks "github.com/XMNBlockchain/core/packages/blocks/blocks/infrastructure"
	validated "github.com/XMNBlockchain/core/packages/blocks/validated/domain"
	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	concrete_hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/infrastructure"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/users/infrastructure"
)

type blockBuilder struct {
	htBuilderFactory hashtrees.HashTreeBuilderFactory
	blk              blocks.Block
	ls               []users.Signature
}

func createBlockBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory) validated.BlockBuilder {
	out := blockBuilder{
		htBuilderFactory: htBuilderFactory,
		blk:              nil,
		ls:               nil,
	}

	return &out
}

// Create initializes the BlockBuilder instance
func (build *blockBuilder) Create() validated.BlockBuilder {
	build.blk = nil
	build.ls = nil
	return build
}

// WithBlock adds a Block to the BlockBuilder instance
func (build *blockBuilder) WithBlock(blk blocks.Block) validated.BlockBuilder {
	build.blk = blk
	return build
}

// WithSignatures adds a leader signatures to the BlockBuilder instance
func (build *blockBuilder) WithSignatures(sigs []users.Signature) validated.BlockBuilder {
	build.ls = sigs
	return build
}

// Now builds a new Block instance
func (build *blockBuilder) Now() (validated.Block, error) {

	if build.blk == nil {
		return nil, errors.New("the block is mandatory in order to build a validated block")
	}

	if build.ls == nil {
		return nil, errors.New("the leader signatures are mandatory in order to build a block instance")
	}

	if len(build.ls) <= 0 {
		return nil, errors.New("the leader signatures cannot be empty in order to build a block instance")
	}

	//add the hashtree hash as the first byte block:
	blks := [][]byte{
		build.blk.GetHashTree().GetHash().Get(),
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

	out := createBlock(ht.(*concrete_hashtrees.HashTree), build.blk.(*concrete_blocks.Block), ls)
	return out, nil
}
