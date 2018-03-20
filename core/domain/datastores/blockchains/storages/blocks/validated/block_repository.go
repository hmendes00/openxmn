package validated

// BlockRepository represents a stored validated block repository
type BlockRepository interface {
	Retrieve(dirPath string) (Block, error)
}
