package tokens

import uuid "github.com/satori/go.uuid"

// Unstake represents an unstake token transaction
type Unstake struct {
	ID      *uuid.UUID `json:"id"`
	StakeID *uuid.UUID `json:"stake_id"`
}

// CreateUnstake creates a new Unstake instance
func CreateUnstake(id *uuid.UUID, stakeID *uuid.UUID) *Unstake {
	out := Unstake{
		ID:      id,
		StakeID: stakeID,
	}

	return &out
}

// GetID returns the ID
func (un *Unstake) GetID() *uuid.UUID {
	return un.ID
}

// GetStakeID returns the stake ID
func (un *Unstake) GetStakeID() *uuid.UUID {
	return un.StakeID
}
