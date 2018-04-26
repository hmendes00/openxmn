package delete

import uuid "github.com/satori/go.uuid"

// Organization represents a delete organization transaction
type Organization struct {
	OrgID *uuid.UUID `json:"organization_id"`
}
