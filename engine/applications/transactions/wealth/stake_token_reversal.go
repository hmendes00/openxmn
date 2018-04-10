package wealth

import uuid "github.com/satori/go.uuid"

// StakeTokenReversal represents a reversal of a stake token transaction
type StakeTokenReversal struct {
	StakeID *uuid.UUID `json:"stake_id"`
}

// CreateStakeTokenReversal creates a new StakeTokenReversal instance
func CreateStakeTokenReversal(stakeID *uuid.UUID) *StakeTokenReversal {
	out := StakeTokenReversal{
		StakeID: stakeID,
	}

	return &out
}

// GetStakeID returns the stakeID
func (st *StakeTokenReversal) GetStakeID() *uuid.UUID {
	return st.StakeID
}
