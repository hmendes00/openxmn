package trees

import (
	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
)

// Branches represents an ordered list of Bfranch
type Branches interface {
	GetHashTree() hashtrees.HashTree
	GetBranches() []Branch
}
