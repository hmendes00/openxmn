package signed

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	stored_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions"
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/users"
)

// Transaction represents a stored signed transaction
type Transaction interface {
	GetMetaData() stored_files.File
	GetSignature() stored_users.Signature
	GetTransaction() stored_transactions.Transaction
}
