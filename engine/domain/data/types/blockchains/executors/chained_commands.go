package executors

import (
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
)

// ChainedCommands represents a chained commands executor
type ChainedCommands interface {
	Execute(chainedCmds commands.ChainedCommands) error
}
