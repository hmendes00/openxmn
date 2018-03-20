package subsets

import (
	sets "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/sets"
)

// SubSet represents a subset of a full set
type SubSet interface {
	GetIndex() int
	GetAmount() int
	GetStoredAmount() int
	GetSet() sets.Set
}
