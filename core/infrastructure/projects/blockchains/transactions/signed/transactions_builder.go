package signed

import (
	"errors"
	"strconv"
	"time"

	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/hashtrees"
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
	signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/transactions/signed"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/metadata"
	uuid "github.com/satori/go.uuid"
)

type transactionsBuilder struct {
	htBuilderFactory  hashtrees.HashTreeBuilderFactory
	metBuilderFactory metadata.MetaDataBuilderFactory
	id                *uuid.UUID
	met               metadata.MetaData
	trs               []signed_transactions.Transaction
	crOn              *time.Time
}

func createTransactionsBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory, metBuilderFactory metadata.MetaDataBuilderFactory) signed_transactions.TransactionsBuilder {
	out := transactionsBuilder{
		htBuilderFactory:  htBuilderFactory,
		metBuilderFactory: metBuilderFactory,
		id:                nil,
		met:               nil,
		trs:               nil,
		crOn:              nil,
	}

	return &out
}

// Create initializes the TransactionsBuilder instance
func (build *transactionsBuilder) Create() signed_transactions.TransactionsBuilder {
	build.id = nil
	build.met = nil
	build.trs = nil
	build.crOn = nil
	return build
}

// WithID adds an ID to the TransactionsBuilder instance
func (build *transactionsBuilder) WithID(id *uuid.UUID) signed_transactions.TransactionsBuilder {
	build.id = id
	return build
}

// WithMetaData adds MetaData to the TransactionsBuilder instance
func (build *transactionsBuilder) WithMetaData(met metadata.MetaData) signed_transactions.TransactionsBuilder {
	build.met = met
	return build
}

// WithTransactions adds []Transaction to the TransactionsBuilder instance
func (build *transactionsBuilder) WithTransactions(trs []signed_transactions.Transaction) signed_transactions.TransactionsBuilder {
	build.trs = trs
	return build
}

// CreatedOn adds creation time to the TransactionsBuilder instance
func (build *transactionsBuilder) CreatedOn(crOn time.Time) signed_transactions.TransactionsBuilder {
	build.crOn = &crOn
	return build
}

// Now builds a new Transactions instance
func (build *transactionsBuilder) Now() (signed_transactions.Transactions, error) {
	if build.trs == nil {
		return nil, errors.New("the []Transaction are mandatory in order to build a Transactions instance")
	}

	if len(build.trs) <= 0 {
		return nil, errors.New("there must be at least 1 transaction in the list in order to build a Transactions instance")
	}

	if build.met == nil {
		if build.id == nil {
			return nil, errors.New("the ID is mandatory in order to build a Transactions instance")
		}

		if build.crOn == nil {
			return nil, errors.New("the creation time is mandatory in order to build a Transactions instance")
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

		met, metErr := build.metBuilderFactory.Create().Create().WithID(build.id).WithHashTree(ht).CreatedOn(*build.crOn).Now()
		if metErr != nil {
			return nil, metErr
		}

		build.met = met
	}

	if build.met == nil {
		return nil, errors.New("the MetaData is mandatory in order to build a Transactions instance")
	}

	trs := []*Transaction{}
	for _, oneTrs := range build.trs {
		trs = append(trs, oneTrs.(*Transaction))
	}

	out := createTransactions(build.met.(*concrete_metadata.MetaData), trs)
	return out, nil
}
