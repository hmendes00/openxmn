package commands

import (
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
)

// InsertBuilderFactory represents a concrete insert builder factory implementation
type InsertBuilderFactory struct {
}

// CreateInsertBuilderFactory creates a new InsertBuilderFactory instance
func CreateInsertBuilderFactory() commands.InsertBuilderFactory {
	out := InsertBuilderFactory{}
	return &out
}

// Create creates a new InsertBuilder instance
func (fac *InsertBuilderFactory) Create() commands.InsertBuilder {
	out := createInsertBuilder()
	return out
}
