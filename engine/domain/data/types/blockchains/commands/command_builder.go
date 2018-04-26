package commands

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
)

// CommandBuilder represents a command builder
type CommandBuilder interface {
	Create() CommandBuilder
	WithMetaData(met metadata.MetaData) CommandBuilder
	WithCommands(cmds []Command) CommandBuilder
	WithInsert(in Insert) CommandBuilder
	WithUpdate(up Update) CommandBuilder
	WithDelete(del Delete) CommandBuilder
	Now() (Command, error)
}
