package domain

// AggregatedSignedTransactionsBuilderFactory represents aggregated aggregated signed transactions builder factory
type AggregatedSignedTransactionsBuilderFactory interface {
	Create() AggregatedSignedTransactionsBuilder
}
