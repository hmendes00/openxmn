package stakes

import (
	"github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	stakes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/stakes"
)

// StakeBuilderFactory represents a concrete StakeBuilderFactory implementation
type StakeBuilderFactory struct {
	metBuilderFactory metadata.BuilderFactory
}

// CreateStakeBuilderFactory creates a new StakeBuilderFactory instance
func CreateStakeBuilderFactory(metBuilderFactory metadata.BuilderFactory) stakes.StakeBuilderFactory {
	out := StakeBuilderFactory{
		metBuilderFactory: metBuilderFactory,
	}

	return &out
}

// Create creates a new StakeBuilder instance
func (fac *StakeBuilderFactory) Create() stakes.StakeBuilder {
	out := createStakeBuilder(fac.metBuilderFactory)
	return out
}
