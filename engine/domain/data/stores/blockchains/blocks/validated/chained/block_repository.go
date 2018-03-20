package chained

// BlockRepository represents a stored chained block repository
type BlockRepository interface {
	Retrieve(dirPath string) (Block, error)
}
