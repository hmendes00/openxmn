package insert

import uuid "github.com/satori/go.uuid"

// Currency represents an insert currency transaction
type Currency struct {
	CurrencyID *uuid.UUID `json:"currency_id"`
	CrOrgID    *uuid.UUID `json:"creator_organization_id"`
	Sym        string     `json:"symbol"`
	Name       string     `json:"name"`
	Desc       string     `json:"description"`
	Amount     float64    `json:"amount"`
}
