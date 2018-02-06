package domain

// TransactionBuilderFactory represents a stored signed TransactionBuilder factory
type TransactionBuilderFactory interface {
	Create() TransactionBuilder
}
