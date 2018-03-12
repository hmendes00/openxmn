package domain

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	stored_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/transactions"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/users"
)

// AtomicTransaction represents a signed atomic transaction
type AtomicTransaction interface {
	GetMetaData() stored_files.File
	GetSignature() stored_users.Signature
	GetTransactions() stored_transactions.Transactions
}
