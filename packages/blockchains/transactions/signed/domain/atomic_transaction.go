package domain

import (
	met "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
)

// AtomicTransaction represents multiple transactions that must be executed all at once, signed by a wallet
type AtomicTransaction interface {
	GetMetaData() met.MetaData
	GetTransactions() transactions.Transactions
	GetSignature() users.Signature
}
