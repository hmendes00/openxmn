package metadata

// Repository represents a metadata repository
type Repository interface {
	Retrieve(dirPath string) (MetaData, error)
}
