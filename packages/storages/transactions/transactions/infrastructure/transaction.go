package infrastructure

import (
	stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/domain"
)

type transaction struct {
	metaData stored_files.File
	chks     stored_chunks.Chunks
}

func createTransaction(metaData stored_files.File, chks stored_chunks.Chunks) stored_transactions.Transaction {
	out := transaction{
		metaData: metaData,
		chks:     chks,
	}

	return &out
}

// GetMetaData returns the metadata file
func (trs *transaction) GetMetaData() stored_files.File {
	return trs.metaData
}

// GetChunks returns the chunks file
func (trs *transaction) GetChunks() stored_chunks.Chunks {
	return trs.chks
}
