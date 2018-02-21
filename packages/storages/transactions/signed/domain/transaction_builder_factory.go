package domain

// TransactionBuilderFactory represents a stored signed transaction builder factory
type TransactionBuilderFactory interface {
	Create() TransactionBuilder
}
