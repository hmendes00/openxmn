package wealth

import uuid "github.com/satori/go.uuid"

// DeleteSafe represents a delete safe transaction
type DeleteSafe struct {
	SafeID *uuid.UUID `json:"safe_id"`
}

// CreateDeleteSafe creates a DeleteSafe instance
func CreateDeleteSafe(safeID *uuid.UUID) *DeleteSafe {
	out := DeleteSafe{
		SafeID: safeID,
	}

	return &out
}

// GetSafeID returns the SafeID
func (del *DeleteSafe) GetSafeID() *uuid.UUID {
	return del.SafeID
}
