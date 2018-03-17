package blocks

// BlockRepository represents a stored block repository
type BlockRepository interface {
	Retrieve(dirPath string) (Block, error)
}
