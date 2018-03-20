package sdks

import (
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	signed "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions/signed"
	servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/servers"
)

// Transactions represents the Transactions SDK
type Transactions interface {
	SaveTrs(serv servers.Server, trs transactions.Transaction) (signed.Transaction, error)
	SaveAtomicTrs(serv servers.Server, trs transactions.Transactions) (signed.AtomicTransaction, error)
}
