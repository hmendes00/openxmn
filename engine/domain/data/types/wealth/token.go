package wealth

import (
	types "github.com/XMNBlockchain/openxmn/engine/domain/data/types"
)

// Token represents a token
type Token interface {
	GetMetaData() types.MetaData
	GetCreator() Entity
	GetSymbol() string
	GetAmount() int
}
