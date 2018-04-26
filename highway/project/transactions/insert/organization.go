package insert

import uuid "github.com/satori/go.uuid"

// Organization represents an insert organization transaction
type Organization struct {
	OrgID      *uuid.UUID `json:"organization_id"`
	CrOrgID    *uuid.UUID `json:"creator_organization_id"`
	SharesID   *uuid.UUID `json:"shares_id"`
	QuotasID   *uuid.UUID `json:"quotas_id"`
	CurrencyID *uuid.UUID `json:"currency_id"`
	Name       string     `json:"name"`
	Desc       string     `json:"description"`
}
