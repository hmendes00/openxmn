package applications

import (
	types "github.com/XMNBlockchain/openxmn/engine/domain/data/types"
	trees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/applications/trees"
)

// Application represents a blockchain application name
type Application interface {
	GetMetaData() types.MetaData
	GetTree() trees.Tree
}
