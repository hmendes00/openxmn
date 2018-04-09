package bills

import (
	bills "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/bills"
)

// PayeeBuilderFactory represents a concrete PayeeBuilderFactory implementation
type PayeeBuilderFactory struct {
}

// CreatePayeeBuilderFactory creates a new PayeeBuilderFactory instance
func CreatePayeeBuilderFactory() bills.PayeeBuilderFactory {
	out := PayeeBuilderFactory{}
	return &out
}

// Create creates a new PayeeBuilder instance
func (fac *PayeeBuilderFactory) Create() bills.PayeeBuilder {
	out := createPayeeBuilder()
	return out
}
