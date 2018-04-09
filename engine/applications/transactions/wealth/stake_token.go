package wealth

import uuid "github.com/satori/go.uuid"

// StakeToken represents a stake token transaction
type StakeToken struct {
	StakeID *uuid.UUID `json:"stake_id"`
	TokenID *uuid.UUID `json:"token_id"`
	Amount  float64    `json:"amount"`
}
