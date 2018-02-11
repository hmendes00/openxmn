package domain

// ObjectRepository represents an object repository
type ObjectRepository interface {
	Retrieve(dirPath string) (Object, error)
	RetrieveAll(dirPath string) ([]Object, error)
}
