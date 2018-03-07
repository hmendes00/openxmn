package infrastructure

import (
	"errors"

	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
	stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/domain"
	stored_users "github.com/XMNBlockchain/core/packages/storages/users/domain"
)

type transactionBuilder struct {
	metaData stored_files.File
	sig      stored_users.Signature
	trs      stored_transactions.Transaction
}

func createTransactionBuilder() stored_signed_transactions.TransactionBuilder {
	out := transactionBuilder{
		metaData: nil,
		sig:      nil,
		trs:      nil,
	}

	return &out
}

// Create initializes the TransactionBuilder instance
func (build *transactionBuilder) Create() stored_signed_transactions.TransactionBuilder {
	build.metaData = nil
	build.sig = nil
	build.trs = nil
	return build
}

// WithMetaData adds a metadata file to the TransactionBuilder instance
func (build *transactionBuilder) WithMetaData(met stored_files.File) stored_signed_transactions.TransactionBuilder {
	build.metaData = met
	return build
}

// WithSignature adds a signature to the TransactionBuilder instance
func (build *transactionBuilder) WithSignature(sig stored_users.Signature) stored_signed_transactions.TransactionBuilder {
	build.sig = sig
	return build
}

// WithTransaction adds a Transaction to the TransactionBuilder instance
func (build *transactionBuilder) WithTransaction(trs stored_transactions.Transaction) stored_signed_transactions.TransactionBuilder {
	build.trs = trs
	return build
}

// Now builds a new Transaction instance
func (build *transactionBuilder) Now() (stored_signed_transactions.Transaction, error) {
	if build.metaData == nil {
		return nil, errors.New("the metadata is mandatory in order to build a Transaction instance")
	}

	if build.sig == nil {
		return nil, errors.New("the signature is mandatory in order to build a Transaction instance")
	}

	if build.trs == nil {
		return nil, errors.New("the trsansaction is mandatory in order to build a Transaction instance")
	}

	out := createTransaction(build.metaData, build.sig, build.trs)
	return out, nil
}
