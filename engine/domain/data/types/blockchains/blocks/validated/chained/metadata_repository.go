package domain

// MetaDataRepository represents a metadata repository
type MetaDataRepository interface {
	Retrieve(dirPath string) (MetaData, error)
}
