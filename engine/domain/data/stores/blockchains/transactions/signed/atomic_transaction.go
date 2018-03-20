package signed

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	stored_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions"
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/users"
)

// AtomicTransaction represents a signed atomic transaction
type AtomicTransaction interface {
	GetMetaData() stored_files.File
	GetSignature() stored_users.Signature
	GetTransactions() stored_transactions.Transactions
}
