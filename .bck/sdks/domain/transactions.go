package domain

import (
	signed "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/domain"
	transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
	servers "github.com/XMNBlockchain/core/packages/servers/domain"
)

// Transactions represents the Transactions SDK
type Transactions interface {
	SaveTrs(serv servers.Server, trs transactions.Transaction) (signed.Transaction, error)
	SaveAtomicTrs(serv servers.Server, trs transactions.Transactions) (signed.AtomicTransaction, error)
}
