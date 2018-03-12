package domain

// TransactionsBuilderFactory represents a TransactionsBuilder factory
type TransactionsBuilderFactory interface {
	Create() TransactionsBuilder
}
