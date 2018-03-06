package domain

// TransactionsBuilderFactory represents a TransactionsBuilderFactory instance
type TransactionsBuilderFactory interface {
	Create() TransactionsBuilder
}
