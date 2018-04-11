package wealth

import uuid "github.com/satori/go.uuid"

// UpdateOrganization  represents a save organization transaction
type UpdateOrganization struct {
	OrgID                     *uuid.UUID `json:"organization_id"`
	TokenID                   *uuid.UUID `json:"token_id"`
	PercentNeededForConcensus float64    `json:"percent_needed_for_concensus"`
}

// CreateUpdateOrganization  creates a new UpdateOrganization  transaction
func CreateUpdateOrganization(orgID *uuid.UUID, tokenID *uuid.UUID, percentNeededForConcensus float64) *UpdateOrganization {
	out := UpdateOrganization{
		OrgID:                     orgID,
		TokenID:                   tokenID,
		PercentNeededForConcensus: percentNeededForConcensus,
	}

	return &out
}

// GetOrganizationID returns the organizationID
func (trs *UpdateOrganization) GetOrganizationID() *uuid.UUID {
	return trs.OrgID
}

// GetTokenID returns the token ID
func (trs *UpdateOrganization) GetTokenID() *uuid.UUID {
	return trs.TokenID
}

// GetPercentNeededForConcensus returns the percentage needed for concensus
func (trs *UpdateOrganization) GetPercentNeededForConcensus() float64 {
	return trs.PercentNeededForConcensus
}
