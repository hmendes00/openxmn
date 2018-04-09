package domain

// SignedBlockRepository represents a SignedBlock repository
type SignedBlockRepository interface {
	Retrieve(dirPath string) (SignedBlock, error)
}
