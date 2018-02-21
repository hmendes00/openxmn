package domain

// BlockRepository represents a Block repository
type BlockRepository interface {
	Retrieve(dirPath string) (Block, error)
}
