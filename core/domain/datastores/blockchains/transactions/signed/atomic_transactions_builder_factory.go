package domain

// AtomicTransactionsBuilderFactory represents an AtomicTransactionsBuilder factory
type AtomicTransactionsBuilderFactory interface {
	Create() AtomicTransactionsBuilder
}
