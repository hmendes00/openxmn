package instances

import uuid "github.com/satori/go.uuid"

// InstantiateApplication represents an instantiate application transaction
type InstantiateApplication struct {
	ID               *uuid.UUID `json:"id"`
	ApplicationID    *uuid.UUID `json:"application_id"`
	TokenID          *uuid.UUID `json:"token_id"`
	MinAmountOfStake float64    `json:"minimum_amount_of_stake"`
	Bounty           float64    `json:"bounty"`
}
