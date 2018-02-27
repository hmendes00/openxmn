package domain

// SignedTransactionsBuilderFactory represents the SignedTransactionsBuilder factory
type SignedTransactionsBuilderFactory interface {
	Create() SignedTransactionsBuilder
}
