package signed

// AtomicTransactionBuilderFactory represents a signed atomic transaction builder factory
type AtomicTransactionBuilderFactory interface {
	Create() AtomicTransactionBuilder
}
