package domain

// TransactionRepository represents a transaction repository
type TransactionRepository interface {
	Retrieve(dirPath string) (Transaction, error)
	RetrieveAll(dirPath string) ([]Transaction, error)
}
