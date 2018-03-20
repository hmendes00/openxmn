package domain

// TransactionsRepository represents a transactions repository
type TransactionsRepository interface {
	Retrieve(dirPath string) (Transactions, error)
}
