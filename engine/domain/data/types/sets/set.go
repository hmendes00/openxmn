package sets

import uuid "github.com/satori/go.uuid"

// Set represents a set of IDs
type Set interface {
	Get() []*uuid.UUID
}
