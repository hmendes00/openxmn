package commands

import (
	bills "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/bills"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
)

// Builder represents a commands builder
type Builder interface {
	Create() Builder
	WithMetaData(met metadata.MetaData) Builder
	WithCommands(cmds []Command) Builder
	WithBill(bi bills.Bill) Builder
	Now() (Commands, error)
}
