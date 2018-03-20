package signed

// TransactionRepository represents a stored transaction repository
type TransactionRepository interface {
	Retrieve(dirPath string) (Transaction, error)
	RetrieveAll(dirPath string) ([]Transaction, error)
}
