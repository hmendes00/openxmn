package signed

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/files"
	stored_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/transactions"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/users"
)

// AtomicTransaction represents a signed atomic transaction
type AtomicTransaction interface {
	GetMetaData() stored_files.File
	GetSignature() stored_users.Signature
	GetTransactions() stored_transactions.Transactions
}
