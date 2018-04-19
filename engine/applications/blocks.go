package main

import (
	"time"

	"github.com/XMNBlockchain/openxmn/engine/applications/agents"
	apis "github.com/XMNBlockchain/openxmn/engine/applications/apis"
	blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks"
	validated_blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks/validated"
	concrete_blocks "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/blocks"
	concrete_blocks_validated "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/blocks/validated"
	concrete_blockchain_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
	concrete_hashtrees "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/hashtrees"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
	"github.com/gorilla/mux"
)

// Blocks represents a blocks application
type Blocks struct {
	api   *apis.Blocks
	agent *agents.ValidateBlock
}

// CreateBlocks creates a new Blocks instance
func CreateBlocks(
	router *mux.Router,
	routePrefix string,
	usersStake map[string]float64,
	signedBlkBufferSize int,
	validatedBlkBufferSize int,
	neededStakerPerBlk float64,
	waitBeforeRemovalTs time.Duration,
) *Blocks {

	//channels:
	newSignedBlock := make(chan blocks.SignedBlock, signedBlkBufferSize)
	newValidatedBlk := make(chan validated_blocks.Block, validatedBlkBufferSize)

	//users related builder factories:
	metaDataBuilderFactory := concrete_metadata.CreateBuilderFactory()
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	userSigsBuilderFactory := concrete_users.CreateSignaturesBuilderFactory(metaDataBuilderFactory)

	//transactions and blocks factories:
	blockMetaDataBuilderFactory := concrete_blockchain_metadata.CreateBuilderFactory()
	signedBlockBuilderFactory := concrete_blocks.CreateSignedBlockBuilderFactory(htBuilderFactory, blockMetaDataBuilderFactory)
	valBlkBuilderFactory := concrete_blocks_validated.CreateBlockBuilderFactory(htBuilderFactory, blockMetaDataBuilderFactory)

	//create the blocks API:
	blocksAPI := apis.CreateBlocks(
		routePrefix,
		router,
		signedBlockBuilderFactory,
		newSignedBlock,
	)

	//create the block agents:
	agent := agents.CreateValidateBlock(
		valBlkBuilderFactory,
		userSigsBuilderFactory,
		neededStakerPerBlk,
		usersStake,
		waitBeforeRemovalTs,
		newSignedBlock,
		newValidatedBlk,
	)

	out := Blocks{
		api:   blocksAPI,
		agent: agent,
	}

	return &out
}
