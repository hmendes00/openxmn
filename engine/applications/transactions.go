package main

import (
	"log"
	"time"

	agents "github.com/XMNBlockchain/openxmn/engine/applications/agents"
	apis "github.com/XMNBlockchain/openxmn/engine/applications/apis"
	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	signed_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions/signed"
	aggregated_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions/signed/aggregated"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	sdks "github.com/XMNBlockchain/openxmn/engine/domain/sdks"
	concrete_cryptography "github.com/XMNBlockchain/openxmn/engine/infrastructure/cryptography"
	concrete_blockchain_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
	concrete_signed_transactions "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/transactions/signed"
	concrete_aggregated_transactions "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/transactions/signed/aggregated"
	concrete_hashtrees "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/hashtrees"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
	concrete_sdks "github.com/XMNBlockchain/openxmn/engine/infrastructure/sdks"
	"github.com/gorilla/mux"
)

// Transactions represents a transactions application
type Transactions struct {
	api              *apis.Transactions
	agent            *agents.PushTransactionsToLeaders
	leadSDK          sdks.Leaders
	servSDK          sdks.Servers
	newAggregatedTrs <-chan aggregated_transactions.Transactions
}

// CreateTransactions creates a new Transactions instance
func CreateTransactions(
	router *mux.Router,
	pk cryptography.PrivateKey,
	user users.User,
	serverSDK sdks.Servers,
	routePrefix string,
	signedTrsBufferSize int,
	atomicTrsBufferSize int,
	aggregatedTrsBufferSize int,
	trsAggregationDelay time.Duration,
) *Transactions {

	//channels:
	newSignedTrs := make(chan signed_transactions.Transaction, signedTrsBufferSize)
	newAtomicSignedTrs := make(chan signed_transactions.AtomicTransaction, atomicTrsBufferSize)
	newAggregatedTrs := make(chan aggregated_transactions.Transactions, aggregatedTrsBufferSize)

	//factories:
	metaDataBuilderFactory := concrete_metadata.CreateBuilderFactory()
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	publicKeyBuilderFactory := concrete_cryptography.CreatePublicKeyBuilderFactory()
	sigBuilderFactory := concrete_cryptography.CreateSignatureBuilderFactory(publicKeyBuilderFactory)
	userSigBuilderFactory := concrete_users.CreateSignatureBuilderFactory(sigBuilderFactory, htBuilderFactory, metaDataBuilderFactory)

	//transactions and blocks factories:
	blockChainMetaDataBuilderFactory := concrete_blockchain_metadata.CreateBuilderFactory()
	signedTrsBuilderFactory := concrete_signed_transactions.CreateTransactionBuilderFactory(htBuilderFactory, blockChainMetaDataBuilderFactory)
	signedTransBuilderFactory := concrete_signed_transactions.CreateTransactionsBuilderFactory(htBuilderFactory, blockChainMetaDataBuilderFactory)
	signedAtomicTransBuilderFactory := concrete_signed_transactions.CreateAtomicTransactionsBuilderFactory(htBuilderFactory, blockChainMetaDataBuilderFactory)
	atomicSignedTrsBuilderFactory := concrete_signed_transactions.CreateAtomicTransactionBuilderFactory(htBuilderFactory, blockChainMetaDataBuilderFactory)
	signedAggregatedTrsBuilderFactory := concrete_aggregated_transactions.CreateTransactionsBuilderFactory(htBuilderFactory, blockChainMetaDataBuilderFactory)

	//create the leader SDK:
	leadSDK := concrete_sdks.CreateLeaders(userSigBuilderFactory, routePrefix, pk, user)

	//create the transaction API:
	transactionsAPI := apis.CreateTransactions(
		routePrefix,
		router,
		signedTrsBuilderFactory,
		atomicSignedTrsBuilderFactory,
		newSignedTrs,
		newAtomicSignedTrs,
	)

	//create the transaction agent:
	trsAgent := agents.CreatePushTransactionsToLeaders(
		signedAggregatedTrsBuilderFactory,
		signedTransBuilderFactory,
		signedAtomicTransBuilderFactory,
		trsAggregationDelay,
		newSignedTrs,
		newAtomicSignedTrs,
		newAggregatedTrs,
	)

	out := Transactions{
		api:              transactionsAPI,
		agent:            trsAgent,
		leadSDK:          leadSDK,
		servSDK:          serverSDK,
		newAggregatedTrs: newAggregatedTrs,
	}

	return &out
}

// Execute execute the transactions application
func (trs *Transactions) Execute() {

	//start the agent:
	go trs.agent.Execute()

	for {
		select {
		case oneAggregatedTrs := <-trs.newAggregatedTrs:

			//retrieve the next leader server:
			lead, leadErr := trs.servSDK.RetrieveNextLeader()
			if leadErr != nil {
				log.Fatalf("there was an error while retrieving the next leader server: %s", leadErr.Error())
			}

			idAsString := oneAggregatedTrs.GetMetaData().GetID().String()
			signedAggregatedTrs, signedAggregatedTrsErr := trs.leadSDK.SaveTrs(lead, oneAggregatedTrs)
			if signedAggregatedTrsErr != nil {
				log.Fatalf("there was an error while saving the aggregated transaction (ID: %s) to server: %s", idAsString, lead.String())
				break
			}

			signedIDAsString := signedAggregatedTrs.GetMetaData().GetID().String()
			log.Printf("successfully pushed aggregated transaction (ID: %s) to server: %s.  Received signed aggregated transaction (ID: %s)", idAsString, lead.String(), signedIDAsString)
			break
		}

	}

}
