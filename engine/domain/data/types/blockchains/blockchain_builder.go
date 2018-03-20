package blockchains

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// BlockchainBuilder represents a blockchain builder
type BlockchainBuilder interface {
	Create() BlockchainBuilder
	WithHeight(height int) BlockchainBuilder
	WithFloorBlockID(floorBlkID *uuid.UUID) BlockchainBuilder
	WithCeilBlockID(ceilBlockID *uuid.UUID) BlockchainBuilder
	CreatedOn(crOn time.Time) BlockchainBuilder
	LastUpdatedOn(lstUpOn time.Time) BlockchainBuilder
	Now() (Blockchain, error)
}
