package bills

import (
	bills "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/bills"
)

// PayerBuilderFactory represents a concrete PayerBuilderFactory implementation
type PayerBuilderFactory struct {
}

// CreatePayerBuilderFactory creates a new PayerBuilderFactory instance
func CreatePayerBuilderFactory() bills.PayerBuilderFactory {
	out := PayerBuilderFactory{}
	return &out
}

// Create creates a new PayerBuilder instance
func (fac *PayerBuilderFactory) Create() bills.PayerBuilder {
	out := createPayerBuilder()
	return out
}
