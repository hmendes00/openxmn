package processors

import (
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions/signed"
)

// Transactions represents a transactions processor
type Transactions interface {
	Process(trs transactions.Transaction) (commands.Command, error)
	AtomicProcess(atomicTrs transactions.AtomicTransaction) (commands.Commands, error)
}
