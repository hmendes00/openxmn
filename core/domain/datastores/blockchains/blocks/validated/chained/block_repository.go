package domain

// BlockRepository represents a block repository
type BlockRepository interface {
	Retrieve(dirPath string) (Block, error)
}
