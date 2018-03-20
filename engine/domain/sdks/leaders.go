package sdks

import (
	aggregated "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/transactions/signed/aggregated"
	servers "github.com/XMNBlockchain/exmachina-network/engine/domain/servers"
)

// Leaders represents the leaders SDK
type Leaders interface {
	SaveTrs(serv servers.Server, trs aggregated.Transactions) (aggregated.SignedTransactions, error)
}
