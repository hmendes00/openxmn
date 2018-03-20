package wealth

import (
	types "github.com/XMNBlockchain/openxmn/engine/domain/data/types"
)

// Stake represents tokens put at stake to an organization, by an entity
type Stake interface {
	GetMetaData() types.MetaData
	ToOrganization() Organization
	FromEntity() Entity
	GetToken() Token
	GetAmount() float64
}
