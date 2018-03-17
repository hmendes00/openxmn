package aggregated

// TransactionsBuilderFactory represents aggregated transactions builder factory
type TransactionsBuilderFactory interface {
	Create() TransactionsBuilder
}
