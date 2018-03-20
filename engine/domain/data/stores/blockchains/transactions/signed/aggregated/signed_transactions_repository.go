package aggregated

// SignedTransactionsRepository represents a stored signed aggregated transactions repository
type SignedTransactionsRepository interface {
	Retrieve(dirPath string) (SignedTransactions, error)
	RetrieveAll(dirPath string) ([]SignedTransactions, error)
}
