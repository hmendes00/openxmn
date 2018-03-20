package signed

// TransactionsRepository represents a stored Transactions repository
type TransactionsRepository interface {
	Retrieve(dirPath string) (Transactions, error)
}
