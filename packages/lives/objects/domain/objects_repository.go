package domain

// ObjectsRepository represents an objects repository
type ObjectsRepository interface {
	Retrieve(dirPath string) (Objects, error)
}
