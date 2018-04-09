package commands

import (
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
)

// ErrorBuilderFactory represents a concrete ErrorBuilderFactory implementation
type ErrorBuilderFactory struct {
}

// CreateErrorBuilderFactory creates a new ErrorBuilderFactory instance
func CreateErrorBuilderFactory() commands.ErrorBuilderFactory {
	out := ErrorBuilderFactory{}
	return &out
}

// Create creates a new ErrorBuilder instance
func (fac *ErrorBuilderFactory) Create() commands.ErrorBuilder {
	out := createErrorBuilder()
	return out
}
