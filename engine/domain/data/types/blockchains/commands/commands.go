package commands

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
)

// Commands represents a list of commands
type Commands interface {
	GetMetaData() metadata.MetaData
	GetCommands() []Command
}
