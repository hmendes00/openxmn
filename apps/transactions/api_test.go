package infrastructure

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"net/url"
	"reflect"
	"testing"

	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	concrete_server "github.com/XMNBlockchain/core/packages/servers/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/users/infrastructure"
	concrete_sdk "github.com/XMNBlockchain/core/sdks/infrastructure"
	uuid "github.com/satori/go.uuid"

	concrete_commons "github.com/XMNBlockchain/core/packages/applications/commons/infrastructure"
	signed_trs "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	concrete_signed_trs "github.com/XMNBlockchain/core/packages/lives/transactions/signed/infrastructure"
	trs "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain"
	concrete_trs "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure"
)

func TestPostTransactions_PostAtomicTransactions_Success(t *testing.T) {

	t.Parallel()

	//variables:
	port := 8080
	userID := uuid.NewV4()
	dbURL, _ := url.Parse(fmt.Sprintf("http://127.0.0.1:%d", port))
	serv, _ := concrete_server.CreateServerBuilderFactory().Create().Create().WithURL(dbURL).Now()

	//generate private key:
	reader := rand.Reader
	bitSize := 4096
	rawPK, _ := rsa.GenerateKey(reader, bitSize)
	pk, _ := cryptography.CreatePrivateKeyBuilderFactory().Create().WithKey(rawPK).Now()

	//create transactions:
	trsList := []trs.Transaction{
		concrete_trs.CreateTransactionForTests(t),
		concrete_trs.CreateTransactionForTests(t),
	}

	//create atomic transactions:
	transactionsList := [][]trs.Transaction{
		{
			concrete_trs.CreateTransactionForTests(t),
			concrete_trs.CreateTransactionForTests(t),
			concrete_trs.CreateTransactionForTests(t),
		},
		{
			concrete_trs.CreateTransactionForTests(t),
		},
		{
			concrete_trs.CreateTransactionForTests(t),
			concrete_trs.CreateTransactionForTests(t),
		},
	}

	//channels:
	newSignedTrs := make(chan signed_trs.Transaction, len(trsList))
	newAtomicSignedTrs := make(chan signed_trs.AtomicTransaction, len(transactionsList))
	defer close(newSignedTrs)
	defer close(newAtomicSignedTrs)

	//factories:
	publicKeyBuilderFactory := concrete_cryptography.CreatePublicKeyBuilderFactory()
	sigBuilderFactory := concrete_cryptography.CreateSignatureBuilderFactory(publicKeyBuilderFactory)
	userBuilderFactory := concrete_users.CreateUserBuilderFactory()
	userSigBuilderFactory := concrete_users.CreateSignatureBuilderFactory(sigBuilderFactory, userBuilderFactory)
	commonSigBuilderFactory := concrete_commons.CreateSignatureBuilderFactory(userSigBuilderFactory)
	signedTrsBuilderFactory := concrete_signed_trs.CreateTransactionBuilderFactory()
	atomicSignedTrsBuilderFactory := concrete_signed_trs.CreateAtomicTransactionBuilderFactory()

	//create application:
	trsApp := CreateAPI(commonSigBuilderFactory, signedTrsBuilderFactory, atomicSignedTrsBuilderFactory, newSignedTrs, newAtomicSignedTrs, port)

	//create SDK:
	trsSDK := concrete_sdk.CreateTransactions(userSigBuilderFactory, pk, &userID)

	//execute:
	go trsApp.Execute()

	//save transactions:
	for index, oneTrs := range trsList {

		signedTrs, signedTrsErr := trsSDK.SaveTrs(serv, oneTrs)
		if signedTrsErr != nil {
			t.Errorf("there was an error while saving the %d transaction: %s", index, signedTrsErr.Error())
		}

		retTrs := signedTrs.GetTrs()
		retUserID := signedTrs.GetSignature().GetUser().GetID()

		if !reflect.DeepEqual(retTrs, oneTrs) {
			t.Errorf("the transaction at index: %d is invalid", index)
		}

		if !reflect.DeepEqual(retUserID, &userID) {
			t.Errorf("the userID in the user signature of the signed transaction is invalid at index: %d.", index)
		}
	}

	//save atomic transactions:
	for index, oneTrsList := range transactionsList {

		signedAtomicTrs, signedAtomicTrsErr := trsSDK.SaveAtomicTrs(serv, oneTrsList)
		if signedAtomicTrsErr != nil {
			t.Errorf("there was an error while saving the %d atomic transactions: %s", index, signedAtomicTrsErr.Error())
		}

		retTrs := signedAtomicTrs.GetTrs()
		retUserID := signedAtomicTrs.GetSignature().GetUser().GetID()

		if len(oneTrsList) != len(retTrs) {
			t.Errorf("the length of the []transaction in the atomic transaction is invalid.  Expected: %d, Returned: %d", len(oneTrsList), len(retTrs))
		}

		if !reflect.DeepEqual(retUserID, &userID) {
			t.Errorf("the userID in the user signature of the atomic signed transaction is invalid at index: %d.", index)
		}
	}

}
