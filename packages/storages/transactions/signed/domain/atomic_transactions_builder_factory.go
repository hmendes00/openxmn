package domain

// AtomicTransactionsBuilderFactory represents stored signed AtomicTransactionsBuilderFactory
type AtomicTransactionsBuilderFactory interface {
	Create() AtomicTransactionsBuilder
}
