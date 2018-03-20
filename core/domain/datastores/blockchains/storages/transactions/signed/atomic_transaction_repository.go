package signed

// AtomicTransactionRepository represents a stored atomic transaction repository
type AtomicTransactionRepository interface {
	Retrieve(dirPath string) (AtomicTransaction, error)
	RetrieveAll(dirPath string) ([]AtomicTransaction, error)
}
