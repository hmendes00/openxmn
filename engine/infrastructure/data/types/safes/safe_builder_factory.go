package safes

import (
	safes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/safes"
)

// SafeBuilderFactory represents a SafeBuilderFactory instance
type SafeBuilderFactory struct {
}

// CreateSafeBuilderFactory creates a new SafeBuilderFactory instance
func CreateSafeBuilderFactory() safes.SafeBuilderFactory {
	out := SafeBuilderFactory{}
	return &out
}

// Create creates a new SafeBuilder instance
func (fac *SafeBuilderFactory) Create() safes.SafeBuilder {
	out := createSafeBuilder()
	return out
}
