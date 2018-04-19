package applications

import uuid "github.com/satori/go.uuid"

// DeleteApplication represents a delete application transactiomn
type DeleteApplication struct {
	ApplicationID *uuid.UUID `json:"application_id"`
}
