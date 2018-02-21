package domain

// TransactionBuilderFactory represents a TransactionBuilder factory
type TransactionBuilderFactory interface {
	Create() TransactionBuilder
}
