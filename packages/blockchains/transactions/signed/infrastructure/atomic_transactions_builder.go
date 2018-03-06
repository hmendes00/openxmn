package infrastructure

import (
	"errors"
	"strconv"
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	signed_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/domain"
	uuid "github.com/satori/go.uuid"
)

type atomicTransactionsBuilder struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.MetaDataBuilderFactory
	id                     *uuid.UUID
	met                    metadata.MetaData
	trs                    []signed_transactions.AtomicTransaction
	crOn                   *time.Time
}

func createAtomicTransactionsBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.MetaDataBuilderFactory) signed_transactions.AtomicTransactionsBuilder {
	out := atomicTransactionsBuilder{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
		id:   nil,
		met:  nil,
		trs:  nil,
		crOn: nil,
	}

	return &out
}

// Create initializes the AtomicTransactionsBuilder instance
func (build *atomicTransactionsBuilder) Create() signed_transactions.AtomicTransactionsBuilder {
	build.id = nil
	build.met = nil
	build.trs = nil
	build.crOn = nil
	return build
}

// WithID adds an ID to the AtomicTransactionsBuilder instance
func (build *atomicTransactionsBuilder) WithID(id *uuid.UUID) signed_transactions.AtomicTransactionsBuilder {
	build.id = id
	return build
}

// WithMetaData adds MetaData to the AtomicTransactionsBuilder instance
func (build *atomicTransactionsBuilder) WithMetaData(met metadata.MetaData) signed_transactions.AtomicTransactionsBuilder {
	build.met = met
	return build
}

// WithTransactions adds []AtomicTransaction to the AtomicTransactionsBuilder instance
func (build *atomicTransactionsBuilder) WithTransactions(trs []signed_transactions.AtomicTransaction) signed_transactions.AtomicTransactionsBuilder {
	build.trs = trs
	return build
}

// CreatedOn adds creation time to the AtomicTransactionsBuilder instance
func (build *atomicTransactionsBuilder) CreatedOn(crOn time.Time) signed_transactions.AtomicTransactionsBuilder {
	build.crOn = &crOn
	return build
}

// Now builds a new AtomicTransactions instance
func (build *atomicTransactionsBuilder) Now() (signed_transactions.AtomicTransactions, error) {
	if build.trs == nil {
		return nil, errors.New("the []AtomicTransaction are mandatory in order to build an AtomicTransactions instance")
	}

	if len(build.trs) <= 0 {
		return nil, errors.New("there must be at least 1 AtomicTransaction in the list in order to build an AtomicTransactions instance")
	}

	if build.met == nil {
		if build.id == nil {
			return nil, errors.New("the ID is mandatory in order to build an AtomicTransactions instance")
		}

		if build.crOn == nil {
			return nil, errors.New("the creation time is mandatory in order to build an AtomicTransactions instance")
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
		return nil, errors.New("the MetaData is mandatory in order to build an AtomicTransactions instance")
	}

	atomicTrs := []*AtomicTransaction{}
	for _, oneTrs := range build.trs {
		atomicTrs = append(atomicTrs, oneTrs.(*AtomicTransaction))
	}

	out := createAtomicTransactions(build.met.(*concrete_metadata.MetaData), atomicTrs)
	return out, nil
}
