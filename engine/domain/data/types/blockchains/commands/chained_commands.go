package commands

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	uuid "github.com/satori/go.uuid"
)

// ChainedCommands represents chained commands
type ChainedCommands interface {
	GetMetaData() metadata.MetaData
	GetCommands() Commands
	GetPreviousID() *uuid.UUID
	GetRootID() *uuid.UUID
}
