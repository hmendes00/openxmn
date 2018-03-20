package signed

// AtomicTransactionsRepository represents a stored atomic transactions repository
type AtomicTransactionsRepository interface {
	Retrieve(dirPath string) (AtomicTransactions, error)
}
