package hashtrees

import (
	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/hashtrees"
)

// Compact represents an HashTree, with only the root hash and the block leaves
type Compact struct {
	h      hashtrees.Hash
	leaves *leaves
}

func createCompactHashTree(h hashtrees.Hash, leaves *leaves) hashtrees.Compact {
	out := Compact{
		h:      h,
		leaves: leaves,
	}

	return &out
}

// GetHash returns the Hash of the Compact hashtree
func (compact *Compact) GetHash() hashtrees.Hash {
	return compact.h
}

// GetLength returns the amount of leaves inside its blockLeaves
func (compact *Compact) GetLength() int {
	return len(compact.leaves.leaves)
}
