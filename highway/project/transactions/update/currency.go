package update

import uuid "github.com/satori/go.uuid"

// Currency represents an update currency transaction
type Currency struct {
	CurrencyID *uuid.UUID `json:"currency_id"`
	Sym        string     `json:"symbol"`
	Name       string     `json:"name"`
	Desc       string     `json:"description"`
}
