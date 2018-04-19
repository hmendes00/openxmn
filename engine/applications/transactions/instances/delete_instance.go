package instances

import uuid "github.com/satori/go.uuid"

// DeleteInstance represents a delete instance transaction
type DeleteInstance struct {
	InstanceID *uuid.UUID `json:"instance_id"`
}
