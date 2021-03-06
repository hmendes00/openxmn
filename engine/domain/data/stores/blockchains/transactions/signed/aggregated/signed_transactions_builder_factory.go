package aggregated

// SignedTransactionsBuilderFactory represents aggregated signed transactions builder factory
type SignedTransactionsBuilderFactory interface {
	Create() SignedTransactionsBuilder
}
