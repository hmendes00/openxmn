package commands

import (
	stored_chained_commands "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/commands"
)

// Service represents a commands service
type Service interface {
	Save(dirPath string, cmd Commands) (stored_chained_commands.Commands, error)
}
