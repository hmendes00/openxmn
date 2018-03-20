package wealth

import (
	types "github.com/XMNBlockchain/openxmn/engine/domain/data/types"
)

// Organization represents an organized group of user
type Organization interface {
	GetMetaData() types.MetaData
	GetAcceptedToken() Token
	GetPercentNeededForConcensus() float64
}
