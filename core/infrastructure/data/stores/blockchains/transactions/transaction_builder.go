package transactions

import (
	"errors"

	stored_chunks "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/chunks"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/files"
	stored_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/transactions"
	concrete_stored_chunks "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/stores/chunks"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/stores/files"
)

type transactionBuilder struct {
	metaData stored_files.File
	chks     stored_chunks.Chunks
}

func createTransactionBuilder() stored_transactions.TransactionBuilder {
	out := transactionBuilder{
		metaData: nil,
		chks:     nil,
	}

	return &out
}

// Create initializes the TransactionBuilder instance
func (build *transactionBuilder) Create() stored_transactions.TransactionBuilder {
	build.metaData = nil
	build.chks = nil
	return build
}

// WithMetaData adds metadata to the TransactionBuilder instance
func (build *transactionBuilder) WithMetaData(met stored_files.File) stored_transactions.TransactionBuilder {
	build.metaData = met
	return build
}

// WithChunks adds chunks to the TransactionBuilder instance
func (build *transactionBuilder) WithChunks(chks stored_chunks.Chunks) stored_transactions.TransactionBuilder {
	build.chks = chks
	return build
}

// Now builds a new Transaction instance
func (build *transactionBuilder) Now() (stored_transactions.Transaction, error) {
	if build.metaData == nil {
		return nil, errors.New("the metadata file is mandatory in order to build a Transaction instance")
	}

	if build.chks == nil {
		return nil, errors.New("the chunk files is mandatory in order to build a Transaction instance")
	}

	out := createTransaction(build.metaData.(*concrete_stored_files.File), build.chks.(*concrete_stored_chunks.Chunks))
	return out, nil
}
