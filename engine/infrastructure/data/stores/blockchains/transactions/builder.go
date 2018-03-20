package transactions

import (
	"errors"

	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	stored_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
)

type builder struct {
	met stored_files.File
	trs []stored_transactions.Transaction
}

func createBuilder() stored_transactions.Builder {
	out := builder{
		met: nil,
		trs: nil,
	}

	return &out
}

// Create initializes the Transactions bulder
func (build *builder) Create() stored_transactions.Builder {
	build.met = nil
	build.trs = nil
	return build
}

// WithMetaData adds MetaData to the Transactions builder
func (build *builder) WithMetaData(met stored_files.File) stored_transactions.Builder {
	build.met = met
	return build
}

// WithTransactions adds Transactions to the Transactions builder
func (build *builder) WithTransactions(trs []stored_transactions.Transaction) stored_transactions.Builder {
	build.trs = trs
	return build
}

// Now builds a new Transactions instance
func (build *builder) Now() (stored_transactions.Transactions, error) {
	if build.met == nil {
		return nil, errors.New("the MetaData is mandatory in order to build a stored Transactions instance")
	}

	if build.trs == nil {
		return nil, errors.New("the Transactions is mandatory in order to build a stored Transactions instance")
	}

	trs := []*Transaction{}
	for _, oneTrs := range build.trs {
		trs = append(trs, oneTrs.(*Transaction))
	}

	out := createTransactions(build.met.(*concrete_stored_files.File), trs)
	return out, nil
}
