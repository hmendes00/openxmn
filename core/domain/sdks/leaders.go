package sdks

import (
	aggregated "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/transactions/signed/aggregated"
	servers "github.com/XMNBlockchain/exmachina-network/core/domain/servers"
)

// Leaders represents the leaders SDK
type Leaders interface {
	SaveTrs(serv servers.Server, trs aggregated.Transactions) (aggregated.SignedTransactions, error)
}
