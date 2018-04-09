package wealth

import uuid "github.com/satori/go.uuid"

// CashSafe represents a cash safe transaction
type CashSafe struct {
	SafeID *uuid.UUID `json:"safe_id"`
}
