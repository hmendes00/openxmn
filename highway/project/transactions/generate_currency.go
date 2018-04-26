package transactions

import (
	uuid "github.com/satori/go.uuid"
)

// GenerateCurrency represents a generate currency transaction
type GenerateCurrency struct {
	ToUserID   *uuid.UUID `json:"to_user_id"`
	ToOrgID    *uuid.UUID `json:"to_organization_id"`
	CurrencyID *uuid.UUID `json:"currency_id"`
	Amount     float64    `json:"amount"`
}
