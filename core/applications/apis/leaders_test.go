package apis

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	users "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/users"
	sdks "github.com/XMNBlockchain/exmachina-network/core/domain/sdks"
	servers "github.com/XMNBlockchain/exmachina-network/core/domain/servers"
	concrete_cryptography "github.com/XMNBlockchain/exmachina-network/core/infrastructure/cryptography/rsa"
	concrete_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/users"
	concrete_sdk "github.com/XMNBlockchain/exmachina-network/core/infrastructure/sdks"
	concrete_server "github.com/XMNBlockchain/exmachina-network/core/infrastructure/servers"
	"github.com/gorilla/mux"

	aggregated_trs "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions/signed/aggregated"
	concrete_aggregated_trs "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/transactions/signed/aggregated"
)

func startLeadersAPI() (*http.Server, []aggregated_trs.Transactions, users.User, sdks.Leaders, servers.Server, chan aggregated_trs.SignedTransactions) {

	//variables:
	port := 8082
	routePrefix := "/transactions"
	router := mux.NewRouter()
	dbURL := fmt.Sprintf("http://127.0.0.1:%d", port)
	serv, _ := concrete_server.CreateServerBuilderFactory().Create().Create().WithURL(dbURL).Now()

	//generate private key:
	reader := rand.Reader
	bitSize := 4096
	rawPK, _ := rsa.GenerateKey(reader, bitSize)
	pk, _ := concrete_cryptography.CreatePrivateKeyBuilderFactory().Create().Create().WithKey(rawPK).Now()

	//create the user:
	rawPubKey := pk.GetKey().PublicKey
	pubKey, _ := concrete_cryptography.CreatePublicKeyBuilderFactory().Create().Create().WithKey(&rawPubKey).Now()
	user := concrete_users.CreateUserUsingProvidedPublicKeyForTests(pubKey)

	//create aggregated transactions:
	aggregatedTrs := []aggregated_trs.Transactions{
		concrete_aggregated_trs.CreateTransactionsForTests(),
		concrete_aggregated_trs.CreateTransactionsForTests(),
		concrete_aggregated_trs.CreateTransactionsForTests(),
		concrete_aggregated_trs.CreateTransactionsForTests(),
	}

	//channels:
	newSignedAggregatedTrs := make(chan aggregated_trs.SignedTransactions, len(aggregatedTrs))

	//factories:
	userSigBuilderFactory := concrete_users.CreateSignatureBuilderFactoryForTests()
	signedAggregatedTrsBuilderFactory := concrete_aggregated_trs.CreateSignedTransactionsBuilderFactoryForTests()

	httpServer := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//create application:
	CreateLeaders(routePrefix, router, signedAggregatedTrsBuilderFactory, newSignedAggregatedTrs)

	//starts the http server:
	go httpServer.ListenAndServe()

	//create SDK:
	sdk := concrete_sdk.CreateLeaders(userSigBuilderFactory, routePrefix, pk, user)

	return httpServer, aggregatedTrs, user, sdk, serv, newSignedAggregatedTrs
}

func TestPostAggregatedTransactions_Success(t *testing.T) {

	//start the API:
	httpServer, aggregatedTrs, user, sdk, serv, newSignedAggregatedTrs := startLeadersAPI()
	defer httpServer.Close()
	defer close(newSignedAggregatedTrs)

	//save aggregated transactions:
	for index, oneAggTrs := range aggregatedTrs {

		signedAggTrs, signedAggTrsErr := sdk.SaveTrs(serv, oneAggTrs)
		if signedAggTrsErr != nil {
			t.Errorf("there was an error while saving the %d aggregated transaction: %s", index, signedAggTrsErr.Error())
		}

		retTrs := signedAggTrs.GetTransactions()
		retUser := signedAggTrs.GetSignature().GetUser()

		if !reflect.DeepEqual(retTrs, oneAggTrs) {
			t.Errorf("the aggregated transaction at index: %d is invalid", index)
		}

		if !reflect.DeepEqual(retUser, user) {
			t.Errorf("the user in the user signature of the aggregated signed transaction is invalid at index: %d.", index)
		}

	}

	//verify that the signed aggregated transactions are also in the channel:
	if len(newSignedAggregatedTrs) != len(aggregatedTrs) {
		t.Errorf("the amount of signed aggregated transactions in the channel is invalid.  Expected: %d, Amount: %d", len(aggregatedTrs), len(newSignedAggregatedTrs))
		return
	}

	//verify in the channel:
	for index, oneAggTrs := range aggregatedTrs {
		signedAggTrs := <-newSignedAggregatedTrs

		retTrs := signedAggTrs.GetTransactions()
		retUser := signedAggTrs.GetSignature().GetUser()

		if !reflect.DeepEqual(retTrs, oneAggTrs) {
			t.Errorf("the aggregated transaction, in the channel, at index: %d is invalid", index)
		}

		if !reflect.DeepEqual(retUser, user) {
			t.Errorf("the user in the user signature of the aggregated signed transaction, in the channel, is invalid at index: %d.", index)
		}
	}

}
