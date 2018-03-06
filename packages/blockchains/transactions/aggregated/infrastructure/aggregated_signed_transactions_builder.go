package infrastructure

import (
	"errors"
	"strconv"
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
	uuid "github.com/satori/go.uuid"
)

type aggregatedSignedTransactionsBuilder struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.MetaDataBuilderFactory
	id                     *uuid.UUID
	met                    metadata.MetaData
	trs                    []aggregated.SignedTransactions
	crOn                   *time.Time
}

func createAggregatedSignedTransactionsBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.MetaDataBuilderFactory) aggregated.AggregatedSignedTransactionsBuilder {
	out := aggregatedSignedTransactionsBuilder{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
		id:   nil,
		met:  nil,
		trs:  nil,
		crOn: nil,
	}

	return &out
}

// Create initializes the AggregatedSignedTransactionsBuilder instance
func (build *aggregatedSignedTransactionsBuilder) Create() aggregated.AggregatedSignedTransactionsBuilder {
	build.id = nil
	build.met = nil
	build.trs = nil
	build.crOn = nil
	return build
}

// WithID adds an ID to the AggregatedSignedTransactionsBuilder instance
func (build *aggregatedSignedTransactionsBuilder) WithID(id *uuid.UUID) aggregated.AggregatedSignedTransactionsBuilder {
	build.id = id
	return build
}

// WithMetaData adds MetaData to the AggregatedSignedTransactionsBuilder instance
func (build *aggregatedSignedTransactionsBuilder) WithMetaData(met metadata.MetaData) aggregated.AggregatedSignedTransactionsBuilder {
	build.met = met
	return build
}

// WithTransactions adds []SignedTransactions to the AggregatedSignedTransactionsBuilder instance
func (build *aggregatedSignedTransactionsBuilder) WithTransactions(trs []aggregated.SignedTransactions) aggregated.AggregatedSignedTransactionsBuilder {
	build.trs = trs
	return build
}

// CreatedOn adds creation time to the AggregatedSignedTransactionsBuilder instance
func (build *aggregatedSignedTransactionsBuilder) CreatedOn(crOn time.Time) aggregated.AggregatedSignedTransactionsBuilder {
	build.crOn = &crOn
	return build
}

// Now builds a new AggregatedSignedTransactions instance
func (build *aggregatedSignedTransactionsBuilder) Now() (aggregated.AggregatedSignedTransactions, error) {
	if build.trs == nil {
		return nil, errors.New("the MetaData is mandatory in order to build an AggregatedSignedTransactions instance ")
	}

	if len(build.trs) <= 0 {
		return nil, errors.New("the amount of SignedTransactions must be greater than 0 in order to build an AggregatedSignedTransactions instance")
	}

	if build.met == nil {
		if build.id == nil {
			return nil, errors.New("the ID is mandatory in order to build an AggregatedSignedTransactions instance")
		}

		if build.crOn == nil {
			return nil, errors.New("the creation time is mandatory in order to build an AggregatedSignedTransactions instance")
		}

		blocks := [][]byte{
			build.id.Bytes(),
			[]byte(strconv.Itoa(int(build.crOn.UnixNano()))),
		}

		for _, oneTrs := range build.trs {
			blocks = append(blocks, oneTrs.GetMetaData().GetHashTree().GetHash().Get())
		}

		ht, htErr := build.htBuilderFactory.Create().Create().WithBlocks(blocks).Now()
		if htErr != nil {
			return nil, htErr
		}

		met, metErr := build.metaDataBuilderFactory.Create().Create().WithID(build.id).WithHashTree(ht).CreatedOn(*build.crOn).Now()
		if metErr != nil {
			return nil, metErr
		}

		build.met = met
	}

	if build.met == nil {
		return nil, errors.New("the MetaData is mandatory in order to build an AggregatedSignedTransactions instance")
	}

	trs := []*SignedTransactions{}
	for _, oneTrs := range build.trs {
		trs = append(trs, oneTrs.(*SignedTransactions))
	}

	out := createAggregatedSignedTransactions(build.met.(*concrete_metadata.MetaData), trs)
	return out, nil
}
