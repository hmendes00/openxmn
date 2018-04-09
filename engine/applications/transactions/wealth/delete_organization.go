package wealth

import uuid "github.com/satori/go.uuid"

// DeleteOrganization represents a delete organization transaction
type DeleteOrganization struct {
	OrgID *uuid.UUID `json:"organization_id"`
}
