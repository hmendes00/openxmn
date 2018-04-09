package wealth

import uuid "github.com/satori/go.uuid"

// SaveOrganization represents a save organization transaction
type SaveOrganization struct {
	ID                        *uuid.UUID `json:"id"`
	TokenID                   *uuid.UUID `json:"token_id"`
	PercentNeededForConcensus float64    `json:"percent_needed_for_concensus"`
}

// CreateSaveOrganization creates a new SaveOrganization transaction
func CreateSaveOrganization(id *uuid.UUID, tokenID *uuid.UUID, percentNeededForConcensus float64) *SaveOrganization {
	out := SaveOrganization{
		ID:                        id,
		TokenID:                   tokenID,
		PercentNeededForConcensus: percentNeededForConcensus,
	}

	return &out
}

// GetID returns the ID
func (trs *SaveOrganization) GetID() *uuid.UUID {
	return trs.ID
}

// GetTokenID returns the token ID
func (trs *SaveOrganization) GetTokenID() *uuid.UUID {
	return trs.TokenID
}

// GetPercentNeededForConcensus returns the percentage needed for concensus
func (trs *SaveOrganization) GetPercentNeededForConcensus() float64 {
	return trs.PercentNeededForConcensus
}
