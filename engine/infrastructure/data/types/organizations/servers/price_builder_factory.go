package servers

import (
	servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/servers"
)

// PriceBuilderFactory represents a concrete PriceBuilderFactory implementation
type PriceBuilderFactory struct {
}

// CreatePriceBuilderFactory creates a new PriceBuilderFactory instance
func CreatePriceBuilderFactory() servers.PriceBuilderFactory {
	out := PriceBuilderFactory{}
	return &out
}

// Create creates a new PriceBuilderFactory instance
func (fac *PriceBuilderFactory) Create() servers.PriceBuilder {
	out := createPriceBuilder()
	return out
}
