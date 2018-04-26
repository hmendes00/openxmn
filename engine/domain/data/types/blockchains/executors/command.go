package executors

import (
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
)

// Command represents a command executor
type Command interface {
	Execute(cmd commands.Command) error
}
