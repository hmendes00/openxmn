package sortedsets

import uuid "github.com/satori/go.uuid"

// SortedSet represents a Set of IDs sorted by a rank
type SortedSet interface {
	Get() []*uuid.UUID
}
