package blockchains

import (
	"math/rand"
	"time"

	blockchains "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains"
	uuid "github.com/satori/go.uuid"
)

// CreateBlockchainForTests creates a Blockchain for tests
func CreateBlockchainForTests() *Blockchain {
	//variables:
	height := (rand.Int() % 50) + 1
	floorID := uuid.NewV4()
	ceilID := uuid.NewV4()
	crOn := time.Now().UTC()
	lastUpOn := crOn.Add(5 * time.Second)

	blkchain := createBlockchain(height, &floorID, &ceilID, crOn, lastUpOn)
	return blkchain.(*Blockchain)
}

// CreateBlockchainBuilderFactoryForTests creates a new BlockchainBuilderFactory for tests
func CreateBlockchainBuilderFactoryForTests() blockchains.BlockchainBuilderFactory {
	out := CreateBlockchainBuilderFactory()
	return out
}
