package complaints

import uuid "github.com/satori/go.uuid"

// Delete reprtesents a delete missing transaction transaction
type Delete struct {
	ID *uuid.UUID `json:"id"`
}
