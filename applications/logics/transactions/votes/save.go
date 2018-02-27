package votes

import uuid "github.com/satori/go.uuid"

// Save represents a save vote on an organization transaction
type Save struct {
	ID     *uuid.UUID `json:"id"`
	OrgID  *uuid.UUID `json:"organization_id"`
	TrsID  *uuid.UUID `json:"transaction_id"`
	IsAppr bool       `json:"is_approved"`
}

// CreateSave represents a new Save instance
func CreateSave(id *uuid.UUID, orgID *uuid.UUID, trsID *uuid.UUID, isApproved bool) *Save {
	out := Save{
		ID:     id,
		OrgID:  orgID,
		TrsID:  trsID,
		IsAppr: isApproved,
	}

	return &out
}

// GetID returns the ID
func (sav *Save) GetID() *uuid.UUID {
	return sav.ID
}

// GetOrganizationID returns the organization ID
func (sav *Save) GetOrganizationID() *uuid.UUID {
	return sav.OrgID
}

// GetTransactionID returns the transaction ID
func (sav *Save) GetTransactionID() *uuid.UUID {
	return sav.TrsID
}

// IsApproved returns true if the transaction is approved, false otherwise
func (sav *Save) IsApproved() bool {
	return sav.IsAppr
}
