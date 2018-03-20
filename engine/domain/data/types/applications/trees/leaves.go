package trees

import (
	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
)

// Leaves represents an ordered list of Leaf
type Leaves interface {
	GetHashTree() hashtrees.HashTree
	GetLeaves() []Leaf
}
