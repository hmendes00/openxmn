package blocks

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Save represents blocks information related to a blockchain
type Save interface {
	GetBlockchainID() *uuid.UUID
	GetNeededStake() float64
	GetDuration() time.Duration
}
