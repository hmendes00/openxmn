package main

import (
	"log"
	"time"

	agents "github.com/XMNBlockchain/openxmn/engine/applications/agents"
	apis "github.com/XMNBlockchain/openxmn/engine/applications/apis"
	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks"
	aggregated_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions/signed/aggregated"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	sdks "github.com/XMNBlockchain/openxmn/engine/domain/sdks"
	concrete_cryptography "github.com/XMNBlockchain/openxmn/engine/infrastructure/cryptography"
	concrete_blocks "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/blocks"
	concrete_blockchain_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
	concrete_aggregated_transactions "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/transactions/signed/aggregated"
	concrete_hashtrees "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/hashtrees"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
	concrete_sdks "github.com/XMNBlockchain/openxmn/engine/infrastructure/sdks"
	"github.com/gorilla/mux"
)

// Leaders represents a leaders application
type Leaders struct {
	api       *apis.Leaders
	agent     *agents.PushAggregatedTransactionsFromLeadersToBlocks
	blockSDK  sdks.Blocks
	serverSDK sdks.Servers
	newBlk    <-chan blocks.Block
}

// CreateLeaders creates a new Leaders instance
func CreateLeaders(
	router *mux.Router,
	pk cryptography.PrivateKey,
	user users.User,
	serverSDK sdks.Servers,
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

	//factories:
	metaDataBuilderFactory := concrete_metadata.CreateBuilderFactory()
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	publicKeyBuilderFactory := concrete_cryptography.CreatePublicKeyBuilderFactory()
	sigBuilderFactory := concrete_cryptography.CreateSignatureBuilderFactory(publicKeyBuilderFactory)
	userSigBuilderFactory := concrete_users.CreateSignatureBuilderFactory(sigBuilderFactory, htBuilderFactory, metaDataBuilderFactory)

	//transactions and blocks factories:
	blockMetaDataBuilderFactory := concrete_blockchain_metadata.CreateBuilderFactory()
	signedAggregatedTrsBuilderFactory := concrete_aggregated_transactions.CreateSignedTransactionsBuilderFactory(htBuilderFactory, blockMetaDataBuilderFactory)
	blkBuilderFactory := concrete_blocks.CreateBlockBuilderFactory(htBuilderFactory, blockMetaDataBuilderFactory)

	//create the block sdk:
	blockSDK := concrete_sdks.CreateBlocks(userSigBuilderFactory, routePrefix, pk, user)

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
		api:       leaderAPI,
		agent:     agent,
		blockSDK:  blockSDK,
		serverSDK: serverSDK,
		newBlk:    newBlk,
	}

	return &out

}

// Execute execute the leaders application
func (lead *Leaders) Execute() {

	//execute the agent:
	go lead.agent.Execute()

	for {
		select {
		case oneBlk := <-lead.newBlk:

			//retrieve the next block server:
			leadServ, leadServErr := lead.serverSDK.RetrieveNextBlock()
			if leadServErr != nil {
				log.Fatalf("there was an error while retrieving the next block server: %s", leadServErr.Error())
			}

			idAsString := oneBlk.GetMetaData().GetID().String()
			signedBlk, signedBlkErr := lead.blockSDK.SaveBlock(leadServ, oneBlk)
			if signedBlkErr != nil {
				log.Fatalf("there was an error while saving the block (ID: %s) to server: %s", idAsString, leadServ.String())
				break
			}

			signedIDAsString := signedBlk.GetMetaData().GetID().String()
			log.Printf("successfully pushed block (ID: %s) to server: %s.  Received signed block (ID: %s)", idAsString, leadServ.String(), signedIDAsString)
			break
		}

	}

}
