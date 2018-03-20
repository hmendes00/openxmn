package transactions

// TransactionRepository represents a stored TransactionRepository
type TransactionRepository interface {
	Retrieve(dirPath string) (Transaction, error)
	RetrieveAll(dirPath string) ([]Transaction, error)
}
