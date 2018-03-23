package domain

import (
	met "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// AtomicTransaction represents multiple transactions that must be executed all at once, signed by a wallet
type AtomicTransaction interface {
	GetMetaData() met.MetaData
	GetTransactions() transactions.Transactions
	GetSignature() users.Signature
}
