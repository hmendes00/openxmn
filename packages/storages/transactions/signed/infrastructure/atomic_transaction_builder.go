package infrastructure

import (
	"errors"

	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
	stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/domain"
)

type atomicTransactionBuilder struct {
	metaData stored_files.File
	sig      stored_files.File
	trs      stored_transactions.Transactions
}

func createAtomicTransactionBuilder() stored_signed_transactions.AtomicTransactionBuilder {
	out := atomicTransactionBuilder{
		metaData: nil,
		sig:      nil,
		trs:      nil,
	}

	return &out
}

// Create initializes the AtomicTransactionBuilder instance
func (build *atomicTransactionBuilder) Create() stored_signed_transactions.AtomicTransactionBuilder {
	build.metaData = nil
	build.sig = nil
	build.trs = nil
	return build
}

// WithMetaData adds a metadata file to the AtomicTransactionBuilder instance
func (build *atomicTransactionBuilder) WithMetaData(met stored_files.File) stored_signed_transactions.AtomicTransactionBuilder {
	build.metaData = met
	return build
}

// WithSignature adds a signature file to the AtomicTransactionBuilder instance
func (build *atomicTransactionBuilder) WithSignature(sig stored_files.File) stored_signed_transactions.AtomicTransactionBuilder {
	build.sig = sig
	return build
}

// WithTransactions adds stored transactions to the AtomicTransactionBuilder instance
func (build *atomicTransactionBuilder) WithTransactions(trs stored_transactions.Transactions) stored_signed_transactions.AtomicTransactionBuilder {
	build.trs = trs
	return build
}

// Now builds a new AtomicTransaction instance
func (build *atomicTransactionBuilder) Now() (stored_signed_transactions.AtomicTransaction, error) {
	if build.metaData == nil {
		return nil, errors.New("the metadata is mandatory in order to build an AtomicTransaction instance")
	}

	if build.sig == nil {
		return nil, errors.New("the signature is mandatory in order to build an AtomicTransaction instance")
	}

	if build.trs == nil {
		return nil, errors.New("the stored transactions is mandatory in order to build an AtomicTransaction instance")
	}

	out := createAtomicTransaction(build.metaData, build.sig, build.trs)
	return out, nil
}
