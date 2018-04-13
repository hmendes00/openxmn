package safes

import (
	safes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/safes"
)

// AmountBuilderFactory represents a concrete AmountBuilderFactory representation
type AmountBuilderFactory struct {
}

// CreateAmountBuilderFactory creates a new AmountBuilderFactory instance
func CreateAmountBuilderFactory() safes.AmountBuilderFactory {
	out := AmountBuilderFactory{}

	return &out
}

// Create creates a new AmountBuilder instance
func (fac *AmountBuilderFactory) Create() safes.AmountBuilder {
	out := createAmountBuilder()
	return out
}
