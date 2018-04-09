package wealth

import uuid "github.com/satori/go.uuid"

// SplitSafe represents a split safe transaction
type SplitSafe struct {
	SafeID *uuid.UUID `json:"safe_id"`
}
