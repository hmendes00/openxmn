package subsets

import (
	sets "github.com/XMNBlockchain/openxmn/engine/domain/data/types/sets"
)

// SubSet represents a subset of a full set
type SubSet interface {
	GetIndex() int
	GetAmount() int
	GetStoredAmount() int
	GetSet() sets.Set
}
