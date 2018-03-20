package domain

// SignedBlockRepository represents a signed block repository
type SignedBlockRepository interface {
	Retrieve(dirPath string) (SignedBlock, error)
}
