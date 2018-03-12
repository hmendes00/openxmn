package domain

import (
	aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
	servers "github.com/XMNBlockchain/core/packages/servers/domain"
)

// Leaders represents the leaders SDK
type Leaders interface {
	SaveTrs(serv servers.Server, trs aggregated.Transactions) (aggregated.SignedTransactions, error)
}
