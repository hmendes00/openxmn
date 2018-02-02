package domain

// AtomicTransactionBuilderFactory represents an AtomicTransactionBuilder factory
type AtomicTransactionBuilderFactory interface {
	Create() AtomicTransactionBuilder
}
