package main

import (
	"time"

	agents "github.com/XMNBlockchain/openxmn/engine/applications/agents"
	apis "github.com/XMNBlockchain/openxmn/engine/applications/apis"
	blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks"
	aggregated_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions/signed/aggregated"
	concrete_blocks "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/blocks"
	concrete_blockchain_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
	concrete_aggregated_transactions "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/transactions/signed/aggregated"
	concrete_hashtrees "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/hashtrees"
	"github.com/gorilla/mux"
)

// Leaders represents a leaders application
type Leaders struct {
	api   *apis.Leaders
	agent *agents.PushAggregatedTransactionsFromLeadersToBlocks
}

// CreateLeaders creates a new Leaders instance
func CreateLeaders(
	router *mux.Router,
	routePrefix string,
	aggregatedSignedTrsBufferSize int,
	aggregatedTrsBufferSize int,
	blkBufferSize int,
	blkTimeDuration time.Duration,
) *Leaders {

	//channels:
	newAggregatedSignedTrs := make(chan aggregated_transactions.SignedTransactions, aggregatedSignedTrsBufferSize)
	newSignedAggrTrs := make(chan aggregated_transactions.SignedTransactions, aggregatedTrsBufferSize)
	newBlk := make(chan blocks.Block, blkBufferSize)

	//transactions and blocks factories:
	metaDataBuilderFactory := concrete_blockchain_metadata.CreateBuilderFactory()
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	signedAggregatedTrsBuilderFactory := concrete_aggregated_transactions.CreateSignedTransactionsBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	blkBuilderFactory := concrete_blocks.CreateBlockBuilderFactory(htBuilderFactory, metaDataBuilderFactory)

	//create the leader API:
	leaderAPI := apis.CreateLeaders(
		routePrefix,
		router,
		signedAggregatedTrsBuilderFactory,
		newAggregatedSignedTrs,
	)

	//create the agent:
	stop := false
	agent := agents.CreatePushAggregatedTransactionsFromLeadersToBlocks(
		blkBuilderFactory,
		blkTimeDuration,
		newSignedAggrTrs,
		newBlk,
		stop,
	)

	out := Leaders{
		api:   leaderAPI,
		agent: agent,
	}

	return &out

}
