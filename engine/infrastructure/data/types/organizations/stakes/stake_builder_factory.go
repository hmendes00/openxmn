package stakes

import (
	stakes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/stakes"
)

// StakeBuilderFactory represents a concrete StakeBuilderFactory implementation
type StakeBuilderFactory struct {
}

// CreateStakeBuilderFactory creates a new StakeBuilderFactory instance
func CreateStakeBuilderFactory() stakes.StakeBuilderFactory {
	out := StakeBuilderFactory{}

	return &out
}

// Create creates a new StakeBuilder instance
func (fac *StakeBuilderFactory) Create() stakes.StakeBuilder {
	out := createStakeBuilder()
	return out
}
