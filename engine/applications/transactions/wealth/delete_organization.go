package wealth

import uuid "github.com/satori/go.uuid"

// DeleteOrganization represents a delete organization transaction
type DeleteOrganization struct {
	OrgID *uuid.UUID `json:"organization_id"`
}

// CreateDeleteOrganization creates a DeleteOrganization instance
func CreateDeleteOrganization(orgID *uuid.UUID) *DeleteOrganization {
	out := DeleteOrganization{
		OrgID: orgID,
	}

	return &out
}

// GetOrganizationID returns the organizationID
func (del *DeleteOrganization) GetOrganizationID() *uuid.UUID {
	return del.OrgID
}
