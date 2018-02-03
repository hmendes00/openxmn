package domain

// Compact represents an HashTree, with only the root hash and the block leaves
type Compact interface {
	GetHash() Hash
	GetLength() int
}
