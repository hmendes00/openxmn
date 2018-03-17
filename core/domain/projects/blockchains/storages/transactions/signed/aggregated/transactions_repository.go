package aggregated

// TransactionsRepository represents a stored aggregated transactions repository
type TransactionsRepository interface {
	Retrieve(dirPath string) (Transactions, error)
}
