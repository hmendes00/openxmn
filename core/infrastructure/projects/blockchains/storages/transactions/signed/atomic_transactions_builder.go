package signed

import (
	"errors"

	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	stored_signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/transactions/signed"
)

type atomicTransactionsBuilder struct {
	met stored_files.File
	trs []stored_signed_transactions.AtomicTransaction
}

func createAtomicTransactionsBuilder() stored_signed_transactions.AtomicTransactionsBuilder {
	out := atomicTransactionsBuilder{
		met: nil,
		trs: nil,
	}

	return &out
}

// Create initializes the AtomicTransactionsBuilder instance
func (build *atomicTransactionsBuilder) Create() stored_signed_transactions.AtomicTransactionsBuilder {
	build.met = nil
	build.trs = nil
	return build
}

// WithMetaData adds MetaData to the AtomicTransactionsBuilder instance
func (build *atomicTransactionsBuilder) WithMetaData(met stored_files.File) stored_signed_transactions.AtomicTransactionsBuilder {
	build.met = met
	return build
}

// WithTransactions adds []AtomicTransaction] to the AtomicTransactionsBuilder instance
func (build *atomicTransactionsBuilder) WithTransactions(trs []stored_signed_transactions.AtomicTransaction) stored_signed_transactions.AtomicTransactionsBuilder {
	build.trs = trs
	return build
}

// Now builds a new AtomicTransactions instance
func (build *atomicTransactionsBuilder) Now() (stored_signed_transactions.AtomicTransactions, error) {
	if build.met == nil {
		return nil, errors.New("the MetaData is mandatory in order to build an AtomicTransactions instance")
	}

	if build.trs == nil {
		return nil, errors.New("the []AtomicTransaction is mandatory in order to build an AtomicTransactions instance")
	}

	out := createAtomicTransactions(build.met, build.trs)
	return out, nil
}
