package commands

import (
	stored_commands "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/commands"
)

// CreateBuilderFactoryForTests creates a new commands BuilderFactory instance for tests
func CreateBuilderFactoryForTests() stored_commands.BuilderFactory {
	out := CreateBuilderFactory()
	return out
}
