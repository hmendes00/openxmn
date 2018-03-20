package sdks

import (
	transactions "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/transactions"
	signed "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/transactions/signed"
	servers "github.com/XMNBlockchain/exmachina-network/core/domain/servers"
)

// Transactions represents the Transactions SDK
type Transactions interface {
	SaveTrs(serv servers.Server, trs transactions.Transaction) (signed.Transaction, error)
	SaveAtomicTrs(serv servers.Server, trs transactions.Transactions) (signed.AtomicTransaction, error)
}
