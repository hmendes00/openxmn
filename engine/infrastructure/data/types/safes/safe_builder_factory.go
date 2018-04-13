package safes

import (
	"github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	safes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/safes"
)

// SafeBuilderFactory represents a SafeBuilderFactory instance
type SafeBuilderFactory struct {
	metaDataBuilderFactory metadata.BuilderFactory
}

// CreateSafeBuilderFactory creates a new SafeBuilderFactory instance
func CreateSafeBuilderFactory(metaDataBuilderFactory metadata.BuilderFactory) safes.SafeBuilderFactory {
	out := SafeBuilderFactory{
		metaDataBuilderFactory: metaDataBuilderFactory,
	}
	return &out
}

// Create creates a new SafeBuilder instance
func (fac *SafeBuilderFactory) Create() safes.SafeBuilder {
	out := createSafeBuilder(fac.metaDataBuilderFactory)
	return out
}
