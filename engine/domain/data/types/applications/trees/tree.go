package trees

import (
	types "github.com/XMNBlockchain/openxmn/engine/domain/data/types"
)

// Tree represents an application tree
type Tree interface {
	GetMetaData() types.MetaData
	GetMaster() Branch
	GetBranches() Branches
}
