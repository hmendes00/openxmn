package domain

// TransactionsBuilderFactory represents the TransactionsBuilder factory
type TransactionsBuilderFactory interface {
	Create() TransactionsBuilder
}
