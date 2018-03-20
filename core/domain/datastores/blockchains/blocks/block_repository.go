package domain

// BlockRepository represents an a block repository
type BlockRepository interface {
	Retrieve(dirPath string) (Block, error)
}
