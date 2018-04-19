package commands

import (
	stored_chained_commands "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/commands"
)

// ChainedCommandsService represents a chained commands service
type ChainedCommandsService interface {
	Save(dirPath string, cmd ChainedCommands) (stored_chained_commands.ChainedCommands, error)
}
