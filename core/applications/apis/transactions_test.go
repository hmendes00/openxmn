package apis

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/users"
	servers "github.com/XMNBlockchain/exmachina-network/core/domain/servers"
	concrete_cryptography "github.com/XMNBlockchain/exmachina-network/core/infrastructure/cryptography/rsa"
	concrete_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/users"
	concrete_sdk "github.com/XMNBlockchain/exmachina-network/core/infrastructure/sdks"
	concrete_server "github.com/XMNBlockchain/exmachina-network/core/infrastructure/servers"
	"github.com/gorilla/mux"

	trs "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/transactions"
	signed_trs "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/transactions/signed"
	"github.com/XMNBlockchain/exmachina-network/core/domain/sdks"
	concrete_trs "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/transactions"
	concrete_signed_trs "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/transactions/signed"

	uuid "github.com/satori/go.uuid"
)

func startTransactionsAPI() (*http.Server, []trs.Transaction, []trs.Transactions, users.User, sdks.Transactions, servers.Server, chan signed_trs.Transaction, chan signed_trs.AtomicTransaction) {

	//factories:
	userSigBuilderFactory := concrete_users.CreateSignatureBuilderFactoryForTests()
	signedTrsBuilderFactory := concrete_signed_trs.CreateTransactionBuilderFactoryForTests()
	atomicSignedTrsBuilderFactory := concrete_signed_trs.CreateAtomicTransactionBuilderFactoryForTests()

	//variables:
	port := 8081
	routePrefix := "/transactions"
	router := mux.NewRouter()
	urlAsString := fmt.Sprintf("http://127.0.0.1:%d", port)
	serv, _ := concrete_server.CreateServerBuilderFactory().Create().Create().WithURL(urlAsString).Now()

	//generate private key:
	reader := rand.Reader
	bitSize := 4096
	rawPK, _ := rsa.GenerateKey(reader, bitSize)
	pk, _ := concrete_cryptography.CreatePrivateKeyBuilderFactory().Create().Create().WithKey(rawPK).Now()

	//create the user:
	rawPubKey := pk.GetKey().PublicKey
	pubKey, _ := concrete_cryptography.CreatePublicKeyBuilderFactory().Create().Create().WithKey(&rawPubKey).Now()
	user := concrete_users.CreateUserUsingProvidedPublicKeyForTests(pubKey)

	//create transactions:
	trsList := []trs.Transaction{
		concrete_trs.CreateTransactionForTests(),
		concrete_trs.CreateTransactionForTests(),
	}

	//create atomic transactions:
	firstID := uuid.NewV4()
	firstCrOn := time.Now().UTC()
	firstMultiTrs, _ := concrete_trs.CreateTransactionsBuilderFactoryForTests().Create().Create().WithID(&firstID).CreatedOn(firstCrOn).WithTransactions([]trs.Transaction{
		concrete_trs.CreateTransactionForTests(),
		concrete_trs.CreateTransactionForTests(),
		concrete_trs.CreateTransactionForTests(),
	}).Now()

	secondID := uuid.NewV4()
	secondCrOn := time.Now().UTC()
	secondMultiTrs, _ := concrete_trs.CreateTransactionsBuilderFactoryForTests().Create().Create().WithID(&secondID).CreatedOn(secondCrOn).WithTransactions([]trs.Transaction{
		concrete_trs.CreateTransactionForTests(),
		concrete_trs.CreateTransactionForTests(),
		concrete_trs.CreateTransactionForTests(),
		concrete_trs.CreateTransactionForTests(),
		concrete_trs.CreateTransactionForTests(),
	}).Now()

	thirdID := uuid.NewV4()
	thirdCrOn := time.Now().UTC()
	thirdMultiTrs, _ := concrete_trs.CreateTransactionsBuilderFactoryForTests().Create().Create().WithID(&thirdID).CreatedOn(thirdCrOn).WithTransactions([]trs.Transaction{
		concrete_trs.CreateTransactionForTests(),
		concrete_trs.CreateTransactionForTests(),
	}).Now()

	transactionsList := []trs.Transactions{
		firstMultiTrs,
		secondMultiTrs,
		thirdMultiTrs,
	}

	//channels:
	newSignedTrs := make(chan signed_trs.Transaction, len(trsList))
	newAtomicSignedTrs := make(chan signed_trs.AtomicTransaction, len(transactionsList))

	httpServer := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//create the api app:
	CreateTransactions(routePrefix, router, signedTrsBuilderFactory, atomicSignedTrsBuilderFactory, newSignedTrs, newAtomicSignedTrs)

	//starts the http server:
	go httpServer.ListenAndServe()

	//create SDK:
	sdk := concrete_sdk.CreateTransactions(userSigBuilderFactory, routePrefix, pk, user)

	return httpServer, trsList, transactionsList, user, sdk, serv, newSignedTrs, newAtomicSignedTrs
}

func TestPostTransactions_PostAtomicTransactions_Success(t *testing.T) {

	//start the API:
	httpServer, trsList, transactionsList, user, sdk, serv, newSignedTrs, newAtomicSignedTrs := startTransactionsAPI()
	defer httpServer.Close()
	defer close(newSignedTrs)
	defer close(newAtomicSignedTrs)

	//save transactions:
	for index, oneTrs := range trsList {

		signedTrs, signedTrsErr := sdk.SaveTrs(serv, oneTrs)
		if signedTrsErr != nil {
			t.Errorf("there was an error while saving the %d transaction: %s", index, signedTrsErr.Error())
		}

		retTrs := signedTrs.GetTransaction()
		retUser := signedTrs.GetSignature().GetUser()

		if !reflect.DeepEqual(retTrs, oneTrs) {
			t.Errorf("the transaction at index: %d is invalid", index)
		}

		if !reflect.DeepEqual(retUser, user) {
			t.Errorf("the user in the user signature of the signed transaction is invalid at index: %d.", index)
		}
	}

	//verify that the transaction is also in the channel:
	if len(newSignedTrs) != len(trsList) {
		t.Errorf("the amount of transactions in the channel is invalid.  Expected: %d, Amount: %d", len(trsList), len(newSignedTrs))
		return
	}

	//verify in the channel:
	for index, oneTrs := range trsList {
		signedTrs := <-newSignedTrs

		retTrs := signedTrs.GetTransaction()
		retUser := signedTrs.GetSignature().GetUser()

		if !reflect.DeepEqual(retTrs, oneTrs) {
			t.Errorf("the transaction, in the channel, at index: %d is invalid", index)
		}

		if !reflect.DeepEqual(retUser, user) {
			t.Errorf("the user in the user signature of the transaction, in the channel, is invalid at index: %d.", index)
		}
	}

	//save atomic transactions:
	for index, oneMultiTrs := range transactionsList {

		signedAtomicTrs, signedAtomicTrsErr := sdk.SaveAtomicTrs(serv, oneMultiTrs)
		if signedAtomicTrsErr != nil {
			t.Errorf("there was an error while saving the %d atomic transactions: %s", index, signedAtomicTrsErr.Error())
		}

		retTrs := signedAtomicTrs.GetTransactions()
		retUser := signedAtomicTrs.GetSignature().GetUser()

		if !reflect.DeepEqual(oneMultiTrs, retTrs) {
			t.Errorf("the Transactions instance inside the AtomicTransactions is invalid")
		}

		if !reflect.DeepEqual(retUser, user) {
			t.Errorf("the user in the user signature of the atomic signed transaction is invalid at index: %d.", index)
		}
	}

	//verify that the transaction is also in the channel:
	if len(newAtomicSignedTrs) != len(transactionsList) {
		t.Errorf("the amount of transactions in the channel is invalid.  Expected: %d, Amount: %d", len(transactionsList), len(newAtomicSignedTrs))
		return
	}

	//verify in the channel:
	for index, oneMultiTrs := range transactionsList {
		signedAtomicTrs := <-newAtomicSignedTrs

		retTrs := signedAtomicTrs.GetTransactions()
		retUser := signedAtomicTrs.GetSignature().GetUser()

		if !reflect.DeepEqual(retTrs, oneMultiTrs) {
			t.Errorf("the atomic transaction, in the channel, at index: %d is invalid", index)
		}

		if !reflect.DeepEqual(retUser, user) {
			t.Errorf("the user in the user signature of the transaction, in the channel, is invalid at index: %d.", index)
		}
	}

}
