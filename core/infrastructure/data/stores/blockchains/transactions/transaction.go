package transactions

import (
	stored_chunks "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/chunks"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/files"
	stored_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/transactions"
	concrete_stored_chunks "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/stores/blockchains/chunks"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/stores/files"
)

// Transaction represents a concrete stored transaction implementation
type Transaction struct {
	MetaData *concrete_stored_files.File    `json:"metadata"`
	Chks     *concrete_stored_chunks.Chunks `json:"chunks"`
}

func createTransaction(metaData *concrete_stored_files.File, chks *concrete_stored_chunks.Chunks) stored_transactions.Transaction {
	out := Transaction{
		MetaData: metaData,
		Chks:     chks,
	}

	return &out
}

// GetMetaData returns the metadata file
func (trs *Transaction) GetMetaData() stored_files.File {
	return trs.MetaData
}

// GetChunks returns the chunks file
func (trs *Transaction) GetChunks() stored_chunks.Chunks {
	return trs.Chks
}
