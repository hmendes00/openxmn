package transactions

// Repository represents a stored transactions repository
type Repository interface {
	Retrieve(dirPath string) (Transactions, error)
}
