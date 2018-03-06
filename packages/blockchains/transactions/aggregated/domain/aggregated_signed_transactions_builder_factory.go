package domain

// AggregatedSignedTransactionsBuilderFactory represents an aggregated signed transactions builder factory
type AggregatedSignedTransactionsBuilderFactory interface {
	Create() AggregatedSignedTransactionsBuilder
}
