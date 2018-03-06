package domain

// TransactionsRepository represents a TransactionsRepository
type TransactionsRepository interface {
	Retrieve(dirPath string) (Transactions, error)
}
