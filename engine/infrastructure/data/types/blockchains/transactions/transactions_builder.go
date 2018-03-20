package transactions

import (
	"errors"
	"strconv"
	"time"

	hashtrees "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/hashtrees"
	metadata "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/metadata"
	transactions "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/transactions"
	trs "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/transactions"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/engine/infrastructure/data/types/blockchains/metadata"
	uuid "github.com/satori/go.uuid"
)

type transactionsBuilder struct {
	htBuilderFactory  hashtrees.HashTreeBuilderFactory
	metBuilderFactory metadata.MetaDataBuilderFactory
	id                *uuid.UUID
	crOn              *time.Time
	met               metadata.MetaData
	trs               []trs.Transaction
}

func createTransactionsBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory, metBuilderFactory metadata.MetaDataBuilderFactory) transactions.TransactionsBuilder {
	out := transactionsBuilder{
		htBuilderFactory:  htBuilderFactory,
		metBuilderFactory: metBuilderFactory,
		id:                nil,
		crOn:              nil,
		met:               nil,
		trs:               nil,
	}

	return &out
}

// Create initializes the TransactionsBuilder instance
func (build *transactionsBuilder) Create() transactions.TransactionsBuilder {
	build.id = nil
	build.crOn = nil
	build.met = nil
	build.trs = nil
	return build
}

// WithID adds an ID to the TransactionsBuilder
func (build *transactionsBuilder) WithID(id *uuid.UUID) transactions.TransactionsBuilder {
	build.id = id
	return build
}

// WithMetaData adds MetaData to the TransactionsBuilder
func (build *transactionsBuilder) WithMetaData(meta metadata.MetaData) transactions.TransactionsBuilder {
	build.met = meta
	return build
}

// WithTransactions adds transactions to the TransactionsBuilder
func (build *transactionsBuilder) WithTransactions(trs []trs.Transaction) transactions.TransactionsBuilder {
	build.trs = trs
	return build
}

// CreatedOn adds creation time to the TransactionsBuilder
func (build *transactionsBuilder) CreatedOn(crOn time.Time) transactions.TransactionsBuilder {
	build.crOn = &crOn
	return build
}

// Now builds a new Transactions instance
func (build *transactionsBuilder) Now() (transactions.Transactions, error) {

	if build.trs == nil {
		return nil, errors.New("the []transaction are mandatory in order to build an Transactions instance")
	}

	if len(build.trs) <= 0 {
		return nil, errors.New("the []transaction cannot be empty")
	}

	if build.met == nil {
		if build.id == nil {
			return nil, errors.New("the ID is mandatory in order to build a Transactions instance")
		}

		if build.crOn == nil {
			return nil, errors.New("the creation time is mandatory in order to build a Transactions instance")
		}

		htBlocks := [][]byte{
			build.id.Bytes(),
			[]byte(strconv.Itoa(int(build.crOn.UnixNano()))),
		}

		for _, oneTrs := range build.trs {
			htBlocks = append(htBlocks, oneTrs.GetMetaData().GetID().Bytes())
		}

		//build the hashtree:
		ht, htErr := build.htBuilderFactory.Create().Create().WithBlocks(htBlocks).Now()
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
		return nil, errors.New("the metadata is mandatory in order to build a Transactions instance")
	}

	trs := []*Transaction{}
	for _, oneTrs := range build.trs {
		trs = append(trs, oneTrs.(*Transaction))
	}

	out := createTransactions(build.met.(*concrete_metadata.MetaData), trs)
	return out, nil
}
