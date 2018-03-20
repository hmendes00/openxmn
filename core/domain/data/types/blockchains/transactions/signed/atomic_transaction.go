package domain

import (
	met "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/metadata"
	transactions "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions"
	users "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/users"
)

// AtomicTransaction represents multiple transactions that must be executed all at once, signed by a wallet
type AtomicTransaction interface {
	GetMetaData() met.MetaData
	GetTransactions() transactions.Transactions
	GetSignature() users.Signature
}
