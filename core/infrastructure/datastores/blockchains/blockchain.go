package blockchains

import (
	"time"

	blockchains "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains"
	uuid "github.com/satori/go.uuid"
)

// Blockchain represents a concrete blockchain implementation
type Blockchain struct {
	Height       int        `json:"height"`
	FloorBlockID *uuid.UUID `json:"floor_block_id"`
	CeilBlockID  *uuid.UUID `json:"ceil_block_id"`
	CrOn         time.Time  `json:"created_on"`
	LastUpOn     time.Time  `json:"last_updated_on"`
}

func createBlockchain(height int, floorBlockID *uuid.UUID, ceilBlockID *uuid.UUID, crOn time.Time, lastUpOn time.Time) blockchains.Blockchain {
	out := Blockchain{
		Height:       height,
		FloorBlockID: floorBlockID,
		CeilBlockID:  ceilBlockID,
		CrOn:         crOn,
		LastUpOn:     lastUpOn,
	}

	return &out
}

// GetHeight returns the height
func (block *Blockchain) GetHeight() int {
	return block.Height
}

// GetFloorBlockID returns the validated floor block ID
func (block *Blockchain) GetFloorBlockID() *uuid.UUID {
	return block.FloorBlockID
}

// GetCeilBlockID returns the chained ceiling block ID
func (block *Blockchain) GetCeilBlockID() *uuid.UUID {
	return block.CeilBlockID
}

// CreatedOn returns the creation time
func (block *Blockchain) CreatedOn() time.Time {
	return block.CrOn
}

// LastUpdatedOn returns the last updated time
func (block *Blockchain) LastUpdatedOn() time.Time {
	return block.LastUpOn
}
