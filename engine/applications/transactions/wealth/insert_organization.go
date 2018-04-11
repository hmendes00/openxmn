package wealth

import uuid "github.com/satori/go.uuid"

// InsertOrganization represents a save organization transaction
type InsertOrganization struct {
	OrgID                     *uuid.UUID `json:"organization_id"`
	TokenID                   *uuid.UUID `json:"token_id"`
	PercentNeededForConcensus float64    `json:"percent_needed_for_concensus"`
}

// CreateInsertOrganization creates a new InsertOrganization transaction
func CreateInsertOrganization(orgID *uuid.UUID, tokenID *uuid.UUID, percentNeededForConcensus float64) *InsertOrganization {
	out := InsertOrganization{
		OrgID:                     orgID,
		TokenID:                   tokenID,
		PercentNeededForConcensus: percentNeededForConcensus,
	}

	return &out
}

// GetOrganizationID returns the organizationID
func (trs *InsertOrganization) GetOrganizationID() *uuid.UUID {
	return trs.OrgID
}

// GetTokenID returns the token ID
func (trs *InsertOrganization) GetTokenID() *uuid.UUID {
	return trs.TokenID
}

// GetPercentNeededForConcensus returns the percentage needed for concensus
func (trs *InsertOrganization) GetPercentNeededForConcensus() float64 {
	return trs.PercentNeededForConcensus
}
