package signed

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/files"
	stored_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/transactions"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/users"
)

// Transaction represents a stored signed transaction
type Transaction interface {
	GetMetaData() stored_files.File
	GetSignature() stored_users.Signature
	GetTransaction() stored_transactions.Transaction
}
