package tokens

import uuid "github.com/satori/go.uuid"

// Unstake represents a return of tokens from a previous stake to an organization
type Unstake interface {
	GetID() *uuid.UUID
	GetStakeID() *uuid.UUID
}
