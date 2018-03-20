package blocks

// SignedBlockRepository represents a stored signed block repository
type SignedBlockRepository interface {
	Retrieve(dirPath string) (SignedBlock, error)
}
