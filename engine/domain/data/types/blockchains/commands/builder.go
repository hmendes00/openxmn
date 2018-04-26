package commands

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
)

// Builder represents a commands builder
type Builder interface {
	Create() Builder
	WithMetaData(met metadata.MetaData) Builder
	WithCommands(cmds []Command) Builder
	Now() (Commands, error)
}
