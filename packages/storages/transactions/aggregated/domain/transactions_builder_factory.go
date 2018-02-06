package domain

// TransactionsBuilderFactory represents stored aggregated transactions builder factory
type TransactionsBuilderFactory interface {
	Create() TransactionsBuilder
}
