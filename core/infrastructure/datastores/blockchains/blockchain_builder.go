package blockchains

import (
	"errors"
	"time"

	blockchains "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains"
	uuid "github.com/satori/go.uuid"
)

type blockchainBuilder struct {
	height       int
	floorBlockID *uuid.UUID
	ceilBlockID  *uuid.UUID
	crOn         *time.Time
	lastUpOn     *time.Time
}

func createBlockchainBuilder() blockchains.BlockchainBuilder {
	out := blockchainBuilder{
		height:       0,
		floorBlockID: nil,
		ceilBlockID:  nil,
		crOn:         nil,
		lastUpOn:     nil,
	}

	return &out
}

// Create initializes the blockchain builder
func (build *blockchainBuilder) Create() blockchains.BlockchainBuilder {
	build.height = 0
	build.floorBlockID = nil
	build.ceilBlockID = nil
	build.crOn = nil
	build.lastUpOn = nil
	return build
}

// WithHeight adds a height to the blockchain builder
func (build *blockchainBuilder) WithHeight(height int) blockchains.BlockchainBuilder {
	build.height = height
	return build
}

// WithFloorBlockID adds a floor block ID to the blockchain builder
func (build *blockchainBuilder) WithFloorBlockID(floorBlkID *uuid.UUID) blockchains.BlockchainBuilder {
	build.floorBlockID = floorBlkID
	return build
}

// WithCeilBlockID adds a ceiling block ID to the blockchain builder
func (build *blockchainBuilder) WithCeilBlockID(ceilBlockID *uuid.UUID) blockchains.BlockchainBuilder {
	build.ceilBlockID = ceilBlockID
	return build
}

// CreatedOn adds a creation time to the blockchain builder
func (build *blockchainBuilder) CreatedOn(crOn time.Time) blockchains.BlockchainBuilder {
	build.crOn = &crOn
	return build
}

// LastUpdatedOn adds a last updated on time to the blockchain builder
func (build *blockchainBuilder) LastUpdatedOn(lstUpOn time.Time) blockchains.BlockchainBuilder {
	build.lastUpOn = &lstUpOn
	return build
}

// Now builds a new Blockchain instance
func (build *blockchainBuilder) Now() (blockchains.Blockchain, error) {
	if build.height == 0 {
		return nil, errors.New("the height is mandatory in order to build a Blockchain instance")
	}

	if build.floorBlockID == nil {
		return nil, errors.New("the floor block ID is mandatory in order to build a Blockchain instance")
	}

	if build.ceilBlockID == nil {
		return nil, errors.New("the ceil block ID is mandatory in order to build a Blockchain instance")
	}

	if build.crOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a Blockchain instance")
	}

	if build.lastUpOn == nil {
		return nil, errors.New("the last updated on time is mandatory in order to build a Blockchain instance")
	}

	if build.lastUpOn.Before(*build.crOn) {
		return nil, errors.New("the last updated on time cannot be berfore the creation time in order to build a Blockchain instance")
	}

	out := createBlockchain(build.height, build.floorBlockID, build.ceilBlockID, *build.crOn, *build.lastUpOn)
	return out, nil
}
