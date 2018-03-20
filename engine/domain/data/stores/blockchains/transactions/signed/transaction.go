package signed

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/files"
	stored_transactions "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/blockchains/transactions"
	stored_users "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/blockchains/users"
)

// Transaction represents a stored signed transaction
type Transaction interface {
	GetMetaData() stored_files.File
	GetSignature() stored_users.Signature
	GetTransaction() stored_transactions.Transaction
}
