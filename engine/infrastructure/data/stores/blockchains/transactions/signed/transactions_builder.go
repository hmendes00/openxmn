package signed

import (
	"errors"

	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	stored_signed_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions/signed"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
)

type transactionsBuilder struct {
	met stored_files.File
	trs []stored_signed_transactions.Transaction
}

func createTransactionsBuilder() stored_signed_transactions.TransactionsBuilder {
	out := transactionsBuilder{
		met: nil,
		trs: nil,
	}

	return &out
}

// Create initializes the TransactionsBuilder instance
func (build *transactionsBuilder) Create() stored_signed_transactions.TransactionsBuilder {
	build.met = nil
	build.trs = nil
	return build
}

// WithMetaData adds MetaData to the TransactionsBuilder instance
func (build *transactionsBuilder) WithMetaData(met stored_files.File) stored_signed_transactions.TransactionsBuilder {
	build.met = met
	return build
}

// WithTransactions adds []Transaction to the TransactionsBuilder instance
func (build *transactionsBuilder) WithTransactions(trs []stored_signed_transactions.Transaction) stored_signed_transactions.TransactionsBuilder {
	build.trs = trs
	return build
}

// Now builds a new Transactions instance
func (build *transactionsBuilder) Now() (stored_signed_transactions.Transactions, error) {
	if build.met == nil {
		return nil, errors.New("the MetaData is mandatory in order to build a Transactions instance")
	}

	if build.trs == nil {
		return nil, errors.New("the []Transaction is mandatory in order to build a Transactions instance")
	}

	trs := []*Transaction{}
	for _, oneTrs := range build.trs {
		trs = append(trs, oneTrs.(*Transaction))
	}

	out := createTransactions(build.met.(*concrete_stored_files.File), trs)
	return out, nil
}
