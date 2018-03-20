package lists

import uuid "github.com/satori/go.uuid"

// List represents a list of IDs
type List interface {
	Get() []*uuid.UUID
}
