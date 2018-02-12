package domain

// HashTreeBuilder represents HashTree builder
type HashTreeBuilder interface {
	Create() HashTreeBuilder
	WithBlocks(blocks [][]byte) HashTreeBuilder
	Now() (HashTree, error)
}
