package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/domain"
	stored_users "github.com/XMNBlockchain/core/packages/storages/users/domain"
)

// Transaction represents a stored signed transaction
type Transaction interface {
	GetMetaData() stored_files.File
	GetSignature() stored_users.Signature
	GetTransaction() stored_transactions.Transaction
}
