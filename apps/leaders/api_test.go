package leaders

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
	aggregated_trs "github.com/XMNBlockchain/core/packages/transactions/aggregated/domain"
	concrete_aggregated_trs "github.com/XMNBlockchain/core/packages/transactions/aggregated/infrastructure"
)

func TestPostAggregatedTransactions_Success(t *testing.T) {

	//variables:
	port := 8082
	userID := uuid.NewV4()
	dbURL, _ := url.Parse(fmt.Sprintf("http://127.0.0.1:%d", port))
	serv, _ := concrete_server.CreateServerBuilderFactory().Create().Create().WithURL(dbURL).Now()

	//generate private key:
	reader := rand.Reader
	bitSize := 4096
	rawPK, _ := rsa.GenerateKey(reader, bitSize)
	pk, _ := cryptography.CreatePrivateKeyBuilderFactory().Create().WithKey(rawPK).Now()

	//create aggregated transactions:
	aggregatedTrs := []aggregated_trs.Transactions{
		concrete_aggregated_trs.CreateTransactionsForTests(t),
		concrete_aggregated_trs.CreateTransactionsForTests(t),
		concrete_aggregated_trs.CreateTransactionsForTests(t),
		concrete_aggregated_trs.CreateTransactionsForTests(t),
	}

	//channels:
	newSignedAggregatedTrs := make(chan aggregated_trs.SignedTransactions, len(aggregatedTrs))

	//factories:
	publicKeyBuilderFactory := concrete_cryptography.CreatePublicKeyBuilderFactory()
	sigBuilderFactory := concrete_cryptography.CreateSignatureBuilderFactory(publicKeyBuilderFactory)
	userBuilderFactory := concrete_users.CreateUserBuilderFactory()
	userSigBuilderFactory := concrete_users.CreateSignatureBuilderFactory(sigBuilderFactory, userBuilderFactory)
	commonSigBuilderFactory := concrete_commons.CreateSignatureBuilderFactory(userSigBuilderFactory)
	signedAggregatedTrsBuilderFactory := concrete_aggregated_trs.CreateSignedTransactionsBuilderFactory(userSigBuilderFactory)

	//create application:
	leadApp := CreateAPI(commonSigBuilderFactory, signedAggregatedTrsBuilderFactory, newSignedAggregatedTrs, port)

	//create sdk:
	leadSDK := concrete_sdk.CreateLeaders(userSigBuilderFactory, pk, &userID)

	//execute:
	go leadApp.Execute()

	//save aggregated transactions:
	for index, oneAggTrs := range aggregatedTrs {

		signedAggTrs, signedAggTrsErr := leadSDK.SaveTrs(serv, oneAggTrs)
		if signedAggTrsErr != nil {
			t.Errorf("there was an error while saving the %d aggregated transaction: %s", index, signedAggTrsErr.Error())
		}

		retTrs := signedAggTrs.GetTrs()
		retUserID := signedAggTrs.GetSignature().GetUser().GetID()

		if !reflect.DeepEqual(retTrs, oneAggTrs) {
			t.Errorf("the aggregated transaction at index: %d is invalid", index)
		}

		if !reflect.DeepEqual(retUserID, &userID) {
			t.Errorf("the userID in the user signature of the aggregated signed transaction is invalid at index: %d.", index)
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

		retTrs := signedAggTrs.GetTrs()
		retUserID := signedAggTrs.GetSignature().GetUser().GetID()

		if !reflect.DeepEqual(retTrs, oneAggTrs) {
			t.Errorf("the aggregated transaction, in the channel, at index: %d is invalid", index)
		}

		if !reflect.DeepEqual(retUserID, &userID) {
			t.Errorf("the userID in the user signature of the aggregated signed transaction, in the channel, is invalid at index: %d.", index)
		}
	}

}
