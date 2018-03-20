package aggregated

import (
	"errors"
	"strconv"
	"time"

	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/hashtrees"
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/metadata"
	signed "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions/signed"
	aggregated "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions/signed/aggregated"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/metadata"
	concrete_signed "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/transactions/signed"
	uuid "github.com/satori/go.uuid"
)

type aggregatedTransactionsBuilder struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.MetaDataBuilderFactory
	id                     *uuid.UUID
	met                    metadata.MetaData
	trs                    signed.Transactions
	atomicTrs              signed.AtomicTransactions
	createdOn              *time.Time
}

func createTransactionsBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.MetaDataBuilderFactory) aggregated.TransactionsBuilder {
	out := aggregatedTransactionsBuilder{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
		id:        nil,
		met:       nil,
		trs:       nil,
		atomicTrs: nil,
		createdOn: nil,
	}

	return &out
}

// Create initializes the TransactionsBuilder instance
func (build *aggregatedTransactionsBuilder) Create() aggregated.TransactionsBuilder {
	build.id = nil
	build.met = nil
	build.trs = nil
	build.atomicTrs = nil
	build.createdOn = nil
	return build
}

// WithID adds an ID to the TransactionsBuilder
func (build *aggregatedTransactionsBuilder) WithID(id *uuid.UUID) aggregated.TransactionsBuilder {
	build.id = id
	return build
}

// WithMetaData adds MetaData to the TransactionsBuilder
func (build *aggregatedTransactionsBuilder) WithMetaData(met metadata.MetaData) aggregated.TransactionsBuilder {
	build.met = met
	return build
}

// WithTransactions adds signed Transactions to the TransactionsBuilder
func (build *aggregatedTransactionsBuilder) WithTransactions(trs signed.Transactions) aggregated.TransactionsBuilder {
	build.trs = trs
	return build
}

// WithAtomicTransactions adds signed AtomicTransactions to the TransactionsBuilder
func (build *aggregatedTransactionsBuilder) WithAtomicTransactions(trs signed.AtomicTransactions) aggregated.TransactionsBuilder {
	build.atomicTrs = trs
	return build
}

// CreatedOn adds a creation time to the TransactionsBuilder
func (build *aggregatedTransactionsBuilder) CreatedOn(ts time.Time) aggregated.TransactionsBuilder {
	build.createdOn = &ts
	return build
}

// Now builds an Transactions instance
func (build *aggregatedTransactionsBuilder) Now() (aggregated.Transactions, error) {

	if build.atomicTrs == nil && build.trs == nil {
		return nil, errors.New("there is no transactions or atomic transactions, therefore the aggregated transactions cannot be built")
	}

	if build.met == nil {
		if build.id == nil {
			return nil, errors.New("the ID is mandatory in order to build a transactions instance")
		}

		if build.createdOn == nil {
			return nil, errors.New("the creation time is mandatory in order to build a transactions instance")
		}

		blocks := [][]byte{
			build.id.Bytes(),
			[]byte(strconv.Itoa(int(build.createdOn.UnixNano()))),
		}

		if build.trs != nil {
			blocks = append(blocks, build.trs.GetMetaData().GetHashTree().GetHash().Get())
		}

		if build.atomicTrs != nil {
			blocks = append(blocks, build.atomicTrs.GetMetaData().GetHashTree().GetHash().Get())
		}

		//build the hashtree:
		ht, htErr := build.htBuilderFactory.Create().Create().WithBlocks(blocks).Now()
		if htErr != nil {
			return nil, htErr
		}

		//build the metadata:
		met, metErr := build.metaDataBuilderFactory.Create().Create().WithID(build.id).WithHashTree(ht).CreatedOn(*build.createdOn).Now()
		if metErr != nil {
			return nil, metErr
		}

		build.met = met
	}

	if build.met == nil {
		return nil, errors.New("the MetaData is mandatory in order to build a Transactions instance")
	}

	if build.trs != nil && build.atomicTrs != nil {
		out := createTransactions(build.met.(*concrete_metadata.MetaData), build.trs.(*concrete_signed.Transactions), build.atomicTrs.(*concrete_signed.AtomicTransactions))
		return out, nil
	}

	if build.trs != nil {
		out := createTransactionsWithTrs(build.met.(*concrete_metadata.MetaData), build.trs.(*concrete_signed.Transactions))
		return out, nil
	}

	out := createTransactionsWithAtomicTrs(build.met.(*concrete_metadata.MetaData), build.atomicTrs.(*concrete_signed.AtomicTransactions))
	return out, nil
}
