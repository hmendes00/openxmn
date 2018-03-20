package transactions

import uuid "github.com/satori/go.uuid"

// SaveTokenByOrganization represents a save token transaction
type SaveTokenByOrganization struct {
	TokenID *uuid.UUID `json:"token_id"`
	Symbol  string     `json:"symbol"`
	Amount  int        `json:"amount"`
}
