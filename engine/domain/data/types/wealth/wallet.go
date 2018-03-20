package wealth

import (
	types "github.com/XMNBlockchain/openxmn/engine/domain/data/types"
)

// Wallet represents a wallet owned by an organization
type Wallet interface {
	GetMetaData() types.MetaData
	GetOwner() Entity
	GetToken() Token
	GetAmount() float64
}
