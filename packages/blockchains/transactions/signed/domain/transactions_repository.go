package domain

// TransactionsRepository represents a TransactionsRepository instance
type TransactionsRepository interface {
	Retrieve(dirPath string) (Transactions, error)
}
