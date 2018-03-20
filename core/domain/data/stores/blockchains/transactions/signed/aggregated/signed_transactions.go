package aggregated

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/files"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/users"
)

// SignedTransactions represents aggregated signed transactions
type SignedTransactions interface {
	GetMetaData() stored_files.File
	GetSignature() stored_users.Signature
	GetTransactions() Transactions
}
