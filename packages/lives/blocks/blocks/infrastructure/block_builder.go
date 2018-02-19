package infrastructure

import (
	"errors"
	"time"

	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/domain"
	concrete_hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/infrastructure"
	aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/domain"
	concrete_aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/infrastructure"
	uuid "github.com/satori/go.uuid"
)

type blockBuilder struct {
	htBuilderFactory hashtrees.HashTreeBuilderFactory
	id               *uuid.UUID
	trs              []aggregated.SignedTransactions
	createdOn        *time.Time
}

func createBlockBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory) blocks.BlockBuilder {
	out := blockBuilder{
		htBuilderFactory: htBuilderFactory,
		id:               nil,
		trs:              nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the BlockBuilder instance
func (build *blockBuilder) Create() blocks.BlockBuilder {
	build.id = nil
	build.trs = nil
	build.createdOn = nil
	return build
}

// WithID adds an ID to the BlockBuilder instance
func (build *blockBuilder) WithID(id *uuid.UUID) blocks.BlockBuilder {
	build.id = id
	return build
}

// WithTransactions adds transactions to the BlockBuilder instance
func (build *blockBuilder) WithTransactions(trs []aggregated.SignedTransactions) blocks.BlockBuilder {
	build.trs = trs
	return build
}

// CreatedOn adds a creation time to the BlockBuilder instance
func (build *blockBuilder) CreatedOn(ts time.Time) blocks.BlockBuilder {
	build.createdOn = &ts
	return build
}

// Now builds a new Block instance
func (build *blockBuilder) Now() (blocks.Block, error) {

	if build.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a Block instance")
	}

	if build.createdOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a Block instance")
	}

	if build.trs == nil {
		return nil, errors.New("the aggregated signed transactions are mandatory in order to build a Block instance")
	}

	if len(build.trs) <= 0 {
		return nil, errors.New("there must be at least 1 aggregate signed transaction instance, none given")
	}

	htBlocks := [][]byte{}
	agregatedSignedTrs := []*concrete_aggregated.SignedTransactions{}
	for _, oneAggSignedTrs := range build.trs {
		agregatedSignedTrs = append(agregatedSignedTrs, oneAggSignedTrs.(*concrete_aggregated.SignedTransactions))
		htBlocks = append(htBlocks, oneAggSignedTrs.GetID().Bytes())
	}

	//build the hashtree:
	ht, htErr := build.htBuilderFactory.Create().Create().WithBlocks(htBlocks).Now()
	if htErr != nil {
		return nil, htErr
	}

	out := createBlock(build.id, ht.(*concrete_hashtrees.HashTree), agregatedSignedTrs, *build.createdOn)
	return out, nil
}
