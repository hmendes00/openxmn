package commands

import (
	bills "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/bills"
)

// Builder represents a commands builder
type Builder interface {
	Create() Builder
	WithCommands(cmds []Command) Builder
	WithBill(bi bills.Bill) Builder
	Now() (Commands, error)
}
