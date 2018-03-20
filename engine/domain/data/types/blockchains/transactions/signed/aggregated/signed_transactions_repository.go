package domain

// SignedTransactionsRepository represents a signed transactions repository
type SignedTransactionsRepository interface {
	Retrieve(dirPath string) (SignedTransactions, error)
	RetrieveAll(dirPath string) ([]SignedTransactions, error)
}
