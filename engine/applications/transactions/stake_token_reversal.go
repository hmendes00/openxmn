package transactions

import uuid "github.com/satori/go.uuid"

// StakeTokenReversal represents a reversal of a stake token transaction
type StakeTokenReversal struct {
	StakeID *uuid.UUID `json:"stake_id"`
}
