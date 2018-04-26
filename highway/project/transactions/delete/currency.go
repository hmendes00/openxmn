package delete

import uuid "github.com/satori/go.uuid"

// Currency represents a delete currency transaction
type Currency struct {
	CurrencyID *uuid.UUID `json:"currency_id"`
}
