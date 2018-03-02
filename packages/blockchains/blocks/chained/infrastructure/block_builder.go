package infrastructure

import (
	"errors"
	"strconv"

	chained "github.com/XMNBlockchain/core/packages/blockchains/blocks/chained/domain"
	validated "github.com/XMNBlockchain/core/packages/blockchains/blocks/validated/domain"
	concrete_validated "github.com/XMNBlockchain/core/packages/blockchains/blocks/validated/infrastructure"
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
)

type blockBuilder struct {
	htBuilderFactory hashtrees.HashTreeBuilderFactory
	met              chained.MetaData
	blk              validated.Block
}

func createBlockBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory) chained.BlockBuilder {
	out := blockBuilder{
		htBuilderFactory: htBuilderFactory,
		met:              nil,
		blk:              nil,
	}

	return &out
}

// Create initializes the BlockBuilder instance
func (build *blockBuilder) Create() chained.BlockBuilder {
	build.met = nil
	build.blk = nil
	return build
}

// WithMetaData adds a MetaData instance to the BlockBuilder
func (build *blockBuilder) WithMetaData(met chained.MetaData) chained.BlockBuilder {
	build.met = met
	return build
}

// WithBlock adds a Block instance to the BlockBuilder
func (build *blockBuilder) WithBlock(blk validated.Block) chained.BlockBuilder {
	build.blk = blk
	return build
}

// Now builds a Block instance
func (build *blockBuilder) Now() (chained.Block, error) {
	if build.met == nil {
		return nil, errors.New("the metadata is mandatory in order to build a Block instance")
	}

	if build.blk == nil {
		return nil, errors.New("the validated block is mandatory in order to build a Block instance")
	}

	blocks := [][]byte{
		build.met.GetID().Bytes(),
		[]byte(strconv.Itoa(build.met.GetIndex())),
		[]byte(strconv.Itoa(build.met.GetPreviousIndex())),
		[]byte(strconv.Itoa(int(build.met.CreatedOn().UnixNano()))),
		build.blk.GetID().Bytes(),
	}

	ht, htErr := build.htBuilderFactory.Create().Create().WithBlocks(blocks).Now()
	if htErr != nil {
		return nil, htErr
	}

	out := createBlock(ht.(*concrete_hashtrees.HashTree), build.met.(*MetaData), build.blk.(*concrete_validated.Block))
	return out, nil
}
