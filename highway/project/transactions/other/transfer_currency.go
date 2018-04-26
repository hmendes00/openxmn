package other

import (
	uuid "github.com/satori/go.uuid"
)

// TransferCurrency represents a transfer currency transaction
type TransferCurrency struct {
	FromOrgID  *uuid.UUID `json:"from_organization_id"`
	ToUserID   *uuid.UUID `json:"to_user_id"`
	ToOrgID    *uuid.UUID `json:"to_organization_id"`
	CurrencyID *uuid.UUID `json:"currency_id"`
	Amount     float64    `json:"amount"`
}
