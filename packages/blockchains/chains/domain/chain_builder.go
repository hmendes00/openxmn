package domain

import (
	"time"

	chained "github.com/XMNBlockchain/core/packages/blockchains/blocks/chained/domain"
	uuid "github.com/satori/go.uuid"
)

// ChainBuilder represents the blockchain builder
type ChainBuilder interface {
	Create() ChainBuilder
	WithID(id *uuid.UUID) ChainBuilder
	WithFloorBlock(floorBlk chained.Block) ChainBuilder
	WithCeilBlock(ceilBlk chained.Block) ChainBuilder
	WithHeight(height int) ChainBuilder
	CreatedOn(crOn time.Time) ChainBuilder
	LastUpdatedOn(lastOn time.Time) ChainBuilder
	Now() (Chain, error)
}
