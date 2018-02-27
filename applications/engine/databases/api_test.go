package databases

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"net/url"
	"reflect"
	"testing"
	"time"

	blocks "github.com/XMNBlockchain/core/packages/blockchains/blocks/blocks/domain"
	concrete_block "github.com/XMNBlockchain/core/packages/blockchains/blocks/blocks/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/blockchains/users/infrastructure"
	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	concrete_server "github.com/XMNBlockchain/core/packages/servers/infrastructure"
	concrete_sdk "github.com/XMNBlockchain/core/sdks/infrastructure"
	uuid "github.com/satori/go.uuid"

	concrete_commons "github.com/XMNBlockchain/core/packages/controllers/signatures/infrastructure"
)

func TestPostBlock_Success(t *testing.T) {

	t.Parallel()

	//variables:
	port := 8086
	userID := uuid.NewV4()
	dbURL, _ := url.Parse(fmt.Sprintf("http://127.0.0.1:%d", port))
	serv, _ := concrete_server.CreateServerBuilderFactory().Create().Create().WithURL(dbURL).Now()

	//generate private key:
	reader := rand.Reader
	bitSize := 4096
	rawPK, _ := rsa.GenerateKey(reader, bitSize)
	pk, _ := cryptography.CreatePrivateKeyBuilderFactory().Create().WithKey(rawPK).Now()

	//create the block:
	blk := concrete_block.CreateBlockForTests(t)

	//channels:
	newSignedBlock := make(chan blocks.SignedBlock, 2)

	//factories:
	publicKeyBuilderFactory := concrete_cryptography.CreatePublicKeyBuilderFactory()
	sigBuilderFactory := concrete_cryptography.CreateSignatureBuilderFactory(publicKeyBuilderFactory)
	userBuilderFactory := concrete_users.CreateUserBuilderFactory()
	userSigBuilderFactory := concrete_users.CreateSignatureBuilderFactory(sigBuilderFactory, userBuilderFactory)
	commonSigBuilderFactory := concrete_commons.CreateSignatureBuilderFactory(userSigBuilderFactory)
	signedBlockBuilderFactory := concrete_block.CreateSignedBlockBuilderFactory()

	//create application:
	dbApp := CreateAPI(commonSigBuilderFactory, signedBlockBuilderFactory, newSignedBlock, port)
	defer dbApp.Stop()

	//execute:
	go dbApp.Execute()

	//sleep some time:
	time.Sleep(time.Second * 2)

	//create sdk:
	dbSDK := concrete_sdk.CreateDatabases(userSigBuilderFactory, pk, &userID)

	//save block:
	signedBlk, signedBlkErr := dbSDK.SaveBlock(serv, blk)
	if signedBlkErr != nil {
		t.Errorf("there was an error while saving a block: %s", signedBlkErr.Error())
		return
	}

	retBlk := signedBlk.GetBlock()
	retUserID := signedBlk.GetSignature().GetUser().GetID()

	if !reflect.DeepEqual(blk, retBlk) {
		t.Errorf("the block inside the signed block is invalid")
		return
	}

	if !reflect.DeepEqual(&userID, retUserID) {
		t.Errorf("the userID inside the signature of the signed block is invalid.  Expected: %s, Returned: %s", userID.String(), retUserID.String())
		return
	}

	//verify the channel:
	assignedSignedBlk := <-newSignedBlock
	if !reflect.DeepEqual(signedBlk, assignedSignedBlk) {
		t.Errorf("the signed block assigned to the channel is invalid")
		return
	}

}
