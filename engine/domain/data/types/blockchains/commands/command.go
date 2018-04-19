package commands

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
)

// Command represents a command
type Command interface {
	GetMetaData() metadata.MetaData
	HasCommands() bool
	GetCommands() Commands
	HasInsert() bool
	GetInsert() Insert
	HasUpdate() bool
	GetUpdate() Update
	HasDelete() bool
	GetDelete() Delete
	HasError() bool
	GetError() Error
}
