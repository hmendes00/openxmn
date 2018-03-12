package blockchains

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Blockchain represents a blockchain
type Blockchain interface {
	GetHeight() int
	GetFloorBlockID() *uuid.UUID
	GetCeilBlockID() *uuid.UUID
	CreatedOn() time.Time
	LastUpdatedOn() time.Time
}
