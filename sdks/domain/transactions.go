package domain

import (
	servers "github.com/XMNBlockchain/core/packages/servers/domain"
	signed "github.com/XMNBlockchain/core/packages/transactions/signed/domain"
	transactions "github.com/XMNBlockchain/core/packages/transactions/transactions/domain"
)

// Transactions represents the Transactions SDK
type Transactions interface {
	SaveTrs(serv servers.Server, trs transactions.Transaction) (signed.Transaction, error)
	SaveAtomicTrs(serv servers.Server, trs []transactions.Transaction) (signed.AtomicTransaction, error)
}
