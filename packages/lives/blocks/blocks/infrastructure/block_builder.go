package infrastructure

import (
	"errors"
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	concrete_hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/infrastructure"
	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/domain"
	concrete_aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/infrastructure"
	"github.com/montanaflynn/stats"
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

	//fetch all the needed karma:
	karmaData := []float64{}
	for _, oneSignedTrs := range build.trs {
		//atomic trs:
		atomicTrs := oneSignedTrs.GetTrs().GetAtomicTrs()
		for _, oneAtomicTrs := range atomicTrs {
			trsList := oneAtomicTrs.GetTrs()
			for _, oneTrs := range trsList {
				karmaData = append(karmaData, float64(oneTrs.GetKarma()))
			}
		}

		//trs:
		trs := oneSignedTrs.GetTrs().GetTrs()
		for _, oneTrs := range trs {
			karmaData = append(karmaData, float64(oneTrs.GetTrs().GetKarma()))
		}
	}

	//make a median with the karma:
	med, medErr := stats.Median(karmaData)
	if medErr != nil {
		return nil, medErr
	}

	//round the median to get the needed karma:
	neededKarma, neededKarmaErr := stats.Round(med, 0)
	if neededKarmaErr != nil {
		return nil, neededKarmaErr
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

	out := createBlock(build.id, ht.(*concrete_hashtrees.HashTree), agregatedSignedTrs, int(neededKarma), *build.createdOn)
	return out, nil
}
