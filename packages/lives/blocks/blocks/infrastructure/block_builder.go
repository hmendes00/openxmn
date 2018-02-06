package infrastructure

import (
	"encoding/json"
	"errors"

	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	hashtree "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	concrete_hashtree "github.com/XMNBlockchain/core/packages/hashtrees/infrastructure"
	aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/domain"
	concrete_aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/infrastructure"
	"github.com/montanaflynn/stats"
)

type blockBuilder struct {
	htBuilderFactory hashtree.HashTreeBuilderFactory
	trs              []aggregated.SignedTransactions
}

func createBlockBuilder(htBuilderFactory hashtree.HashTreeBuilderFactory) blocks.BlockBuilder {
	out := blockBuilder{
		htBuilderFactory: htBuilderFactory,
		trs:              nil,
	}

	return &out
}

// Create initializes the BlockBuilder instance
func (build *blockBuilder) Create() blocks.BlockBuilder {
	build.trs = nil
	return build
}

// WithTransactions adds transactions to the BlockBuilder instance
func (build *blockBuilder) WithTransactions(trs []aggregated.SignedTransactions) blocks.BlockBuilder {
	build.trs = trs
	return build
}

// Now builds a new Block instance
func (build *blockBuilder) Now() (blocks.Block, error) {

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

	//convert the trs to json:
	js, jsErr := json.Marshal(build.trs)
	if jsErr != nil {
		return nil, jsErr
	}

	//build the hashtree:
	ht, htErr := build.htBuilderFactory.Create().Create().WithJSON(js).Now()
	if htErr != nil {
		return nil, htErr
	}

	agregatedSignedTrs := []*concrete_aggregated.SignedTransactions{}
	for _, oneAggSignedTrs := range build.trs {
		agregatedSignedTrs = append(agregatedSignedTrs, oneAggSignedTrs.(*concrete_aggregated.SignedTransactions))
	}

	out := createBlock(ht.(*concrete_hashtree.HashTree), agregatedSignedTrs, int(neededKarma))
	return out, nil
}
