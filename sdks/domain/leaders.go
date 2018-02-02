package domain

import (
	servers "github.com/XMNBlockchain/core/packages/servers/domain"
	aggregated "github.com/XMNBlockchain/core/packages/transactions/aggregated/domain"
)

// Leaders represents the leaders SDK
type Leaders interface {
	SaveTrs(serv servers.Server, trs aggregated.Transactions) (aggregated.SignedTransactions, error)
}
