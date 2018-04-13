package aggregated

import (
	"errors"

	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	stored_aggregated_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions/signed/aggregated"
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/users"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
	concrete_stored_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/users"
)

type signedTransactionsBuilder struct {
	metaData stored_files.File
	sig      stored_users.Signature
	trs      stored_aggregated_transactions.Transactions
}

func createSignedTransactionsBuilder() stored_aggregated_transactions.SignedTransactionsBuilder {
	out := signedTransactionsBuilder{
		metaData: nil,
		sig:      nil,
		trs:      nil,
	}

	return &out
}

// Create initializes the SignedTransactionsBuilder instance
func (build *signedTransactionsBuilder) Create() stored_aggregated_transactions.SignedTransactionsBuilder {
	build.metaData = nil
	build.sig = nil
	build.trs = nil
	return build
}

// WithMetaData adds a metadata file to the SignedTransactionsBuilder instance
func (build *signedTransactionsBuilder) WithMetaData(met stored_files.File) stored_aggregated_transactions.SignedTransactionsBuilder {
	build.metaData = met
	return build
}

// WithSignature adds a signature to the SignedTransactionsBuilder instance
func (build *signedTransactionsBuilder) WithSignature(sig stored_users.Signature) stored_aggregated_transactions.SignedTransactionsBuilder {
	build.sig = sig
	return build
}

// WithTransactions adds stored transactions to the SignedTransactionsBuilder instance
func (build *signedTransactionsBuilder) WithTransactions(trs stored_aggregated_transactions.Transactions) stored_aggregated_transactions.SignedTransactionsBuilder {
	build.trs = trs
	return build
}

// Now builds a new SignedTransactions instance
func (build *signedTransactionsBuilder) Now() (stored_aggregated_transactions.SignedTransactions, error) {
	if build.metaData == nil {
		return nil, errors.New("the metadata file is mandatory in order to build a SignedTransactions instance")
	}

	if build.sig == nil {
		return nil, errors.New("the signature file is mandatory in order to build a SignedTransactions instance")
	}

	if build.trs == nil {
		return nil, errors.New("the stored transactions are mandatory in order to build a SignedTransactions instance")
	}

	out := createSignedTransactions(build.metaData.(*concrete_stored_files.File), build.sig.(*concrete_stored_users.Signature), build.trs.(*Transactions))
	return out, nil
}
