package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/domain"
	stored_users "github.com/XMNBlockchain/core/packages/storages/users/domain"
)

// AtomicTransaction represents a signed atomic transaction
type AtomicTransaction interface {
	GetMetaData() stored_files.File
	GetSignature() stored_users.Signature
	GetTransactions() stored_transactions.Transactions
}
