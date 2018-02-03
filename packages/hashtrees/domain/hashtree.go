package domain

// HashTree represents an HashTree
type HashTree interface {
	GetHeight() int
	GetLength() int
	GetHash() Hash
	Compact() Compact
}
