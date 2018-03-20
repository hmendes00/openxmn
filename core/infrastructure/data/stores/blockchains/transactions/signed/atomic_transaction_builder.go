package signed

import (
	"errors"

	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/files"
	stored_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/transactions"
	stored_signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/transactions/signed"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/users"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/stores/blockchains/files"
	conrete_stored_transactions "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/stores/blockchains/transactions"
	conrete_stored_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/stores/blockchains/users"
)

type atomicTransactionBuilder struct {
	metaData stored_files.File
	sig      stored_users.Signature
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

// WithSignature adds a signature to the AtomicTransactionBuilder instance
func (build *atomicTransactionBuilder) WithSignature(sig stored_users.Signature) stored_signed_transactions.AtomicTransactionBuilder {
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

	out := createAtomicTransaction(build.metaData.(*concrete_stored_files.File), build.sig.(*conrete_stored_users.Signature), build.trs.(*conrete_stored_transactions.Transactions))
	return out, nil
}
