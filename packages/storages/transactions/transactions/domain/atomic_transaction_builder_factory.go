package domain

// AtomicTransactionBuilderFactory represents a stored atomic signed TransactionBuilder factory
type AtomicTransactionBuilderFactory interface {
	Create() AtomicTransactionBuilder
}
