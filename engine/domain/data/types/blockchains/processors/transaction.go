package processors

import (
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// Transaction represents a transaction processor
type Transaction interface {
	Process(trs transactions.Transaction, usr users.User) (commands.Command, error)
}
