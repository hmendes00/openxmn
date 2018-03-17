package transactions

// TransactionBuilderFactory represents a TransactionBuilder factory
type TransactionBuilderFactory interface {
	Create() TransactionBuilder
}
