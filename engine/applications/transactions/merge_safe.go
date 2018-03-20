package transactions

import uuid "github.com/satori/go.uuid"

// MergeSafe represents a merge safe transaction
type MergeSafe struct {
	SafeID *uuid.UUID `json:"safe_id"`
}
