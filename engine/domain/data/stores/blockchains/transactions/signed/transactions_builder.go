package signed

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
)

// TransactionsBuilder represents stored signed TransactionsBuilder
type TransactionsBuilder interface {
	Create() TransactionsBuilder
	WithMetaData(met stored_files.File) TransactionsBuilder
	WithTransactions(trs []Transaction) TransactionsBuilder
	Now() (Transactions, error)
}
