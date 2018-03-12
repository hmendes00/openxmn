package domain

// TransactionsBuilderFactory represents stored signed TransactionsBuilderFactory
type TransactionsBuilderFactory interface {
	Create() TransactionsBuilder
}
