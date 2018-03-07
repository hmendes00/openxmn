package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_users "github.com/XMNBlockchain/core/packages/storages/users/domain"
)

// SignedTransactions represents aggregated signed transactions
type SignedTransactions interface {
	GetMetaData() stored_files.File
	GetSignature() stored_users.Signature
	GetTransactions() Transactions
}
