package infrastructure

import (
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
)

// SignedTransactionsBuilderFactory represents the concrete SignedTransactionsBuilder factory
type SignedTransactionsBuilderFactory struct {
	sigBuilderFactory users.SignatureBuilderFactory
}

// CreateSignedTransactionsBuilderFactory creates a new SignedTransactionsBuilderFactory instance
func CreateSignedTransactionsBuilderFactory(sigBuilderFactory users.SignatureBuilderFactory) aggregated.SignedTransactionsBuilderFactory {
	out := SignedTransactionsBuilderFactory{
		sigBuilderFactory: sigBuilderFactory,
	}
	return &out
}

// Create creates a new SignedTransactionsBuilder instance
func (fac *SignedTransactionsBuilderFactory) Create() aggregated.SignedTransactionsBuilder {
	out := createSignedTransactionsBuilder(fac.sigBuilderFactory)
	return out
}
