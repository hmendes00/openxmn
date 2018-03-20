package aggregated

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/users"
)

// SignedTransactions represents aggregated signed transactions
type SignedTransactions interface {
	GetMetaData() stored_files.File
	GetSignature() stored_users.Signature
	GetTransactions() Transactions
}
