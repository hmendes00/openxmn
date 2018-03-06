package domain

// AtomicTransactionsRepository represents an AtomicTransactionsRepository instance
type AtomicTransactionsRepository interface {
	Retrieve(dirPath string) (AtomicTransactions, error)
}
