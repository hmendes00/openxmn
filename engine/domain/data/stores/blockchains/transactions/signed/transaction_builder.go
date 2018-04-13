package signed

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	stored_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions"
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/users"
)

// TransactionBuilder represents a stored signed transaction builder
type TransactionBuilder interface {
	Create() TransactionBuilder
	WithMetaData(met stored_files.File) TransactionBuilder
	WithSignature(sig stored_users.Signature) TransactionBuilder
	WithTransaction(trs stored_transactions.Transaction) TransactionBuilder
	Now() (Transaction, error)
}
