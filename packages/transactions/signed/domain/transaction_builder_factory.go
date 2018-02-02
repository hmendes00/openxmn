package domain

// TransactionBuilderFactory represents a signed TransactionBuilder factory
type TransactionBuilderFactory interface {
	Create() TransactionBuilder
}
