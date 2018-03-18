package apis

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/blocks"
	users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/users"
	sdks "github.com/XMNBlockchain/exmachina-network/core/domain/sdks"
	servers "github.com/XMNBlockchain/exmachina-network/core/domain/servers"
	concrete_cryptography "github.com/XMNBlockchain/exmachina-network/core/infrastructure/cryptography/rsa"
	concrete_block "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/blocks"
	concrete_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/users"
	concrete_sdk "github.com/XMNBlockchain/exmachina-network/core/infrastructure/sdks"
	concrete_server "github.com/XMNBlockchain/exmachina-network/core/infrastructure/servers"
	"github.com/gorilla/mux"
)

func startBlocksAPI() (*http.Server, blocks.Block, users.User, sdks.Blocks, servers.Server, chan blocks.SignedBlock) {

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

	//create the block:
	blk := concrete_block.CreateBlockForTests()

	//channels:
	newSignedBlock := make(chan blocks.SignedBlock, 2)

	//factories:
	userSigBuilderFactory := concrete_users.CreateSignatureBuilderFactoryForTests()
	signedBlockBuilderFactory := concrete_block.CreateSignedBlockBuilderFactoryForTests()

	httpServer := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//create application:
	CreateBlocks(routePrefix, router, signedBlockBuilderFactory, newSignedBlock)

	//starts the http server:
	go httpServer.ListenAndServe()

	//create SDK:
	sdk := concrete_sdk.CreateBlocks(userSigBuilderFactory, routePrefix, pk, user)

	return httpServer, blk, user, sdk, serv, newSignedBlock
}

func TestPostBlock_Success(t *testing.T) {

	//start the API:
	httpServer, blk, user, sdk, serv, newSignedBlock := startBlocksAPI()
	defer httpServer.Close()
	defer close(newSignedBlock)

	//save block:
	signedBlk, signedBlkErr := sdk.SaveBlock(serv, blk)
	if signedBlkErr != nil {
		t.Errorf("there was an error while saving a block: %s", signedBlkErr.Error())
		return
	}

	retBlk := signedBlk.GetBlock()
	retUser := signedBlk.GetSignature().GetUser()

	if !reflect.DeepEqual(blk, retBlk) {
		t.Errorf("the block inside the signed block is invalid")
		return
	}

	if !reflect.DeepEqual(user, retUser) {
		t.Errorf("the user inside the signature of the signed block is invalid")
		return
	}

	//verify the channel:
	assignedSignedBlk := <-newSignedBlock
	if !reflect.DeepEqual(signedBlk, assignedSignedBlk) {
		t.Errorf("the signed block assigned to the channel is invalid")
		return
	}

}
