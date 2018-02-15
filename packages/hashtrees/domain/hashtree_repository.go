package domain

// HashTreeRepository represents a n HashTree repository
type HashTreeRepository interface {
	Retrieve(dirpath string) (HashTree, error)
}
