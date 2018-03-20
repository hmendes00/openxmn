package agents

import (
	"crypto/rand"
	"crypto/rsa"
	mathrand "math/rand"
	"reflect"
	"testing"
	"time"

	cryptography "github.com/XMNBlockchain/exmachina-network/core/domain/cryptography"
	blocks "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/blocks"
	validated_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/blocks/validated"
	users "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/users"
	concrete_cryptography "github.com/XMNBlockchain/exmachina-network/core/infrastructure/cryptography/rsa"
	concrete_block "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/blocks"
	concrete_validated_block "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/blocks/validated"
	concrete_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/users"
	uuid "github.com/satori/go.uuid"
)

func TestValidateBlock_Success(t *testing.T) {

	//create users:
	generateUserAndPKFn := func() (users.User, cryptography.PrivateKey) {
		//generate private key:
		reader := rand.Reader
		bitSize := 4096
		rawPK, _ := rsa.GenerateKey(reader, bitSize)
		pk, _ := concrete_cryptography.CreatePrivateKeyBuilderFactory().Create().Create().WithKey(rawPK).Now()

		//create the user:
		rawPubKey := pk.GetKey().PublicKey
		pubKey, _ := concrete_cryptography.CreatePublicKeyBuilderFactory().Create().Create().WithKey(&rawPubKey).Now()
		user := concrete_users.CreateUserUsingProvidedPublicKeyForTests(pubKey)
		return user, pk
	}

	//gets the validate block from the passed chan:
	getValidatedBlockFn := func(newValidatedBlk <-chan validated_blocks.Block) validated_blocks.Block {
		select {
		case oneValidatedBlock := <-newValidatedBlk:
			return oneValidatedBlock
		default:
			return nil
		}
	}

	//validate block fn:
	verifyValidatedBlockFn := func(validatedBlk validated_blocks.Block, blk blocks.Block, firstUser users.User, secondUser users.User, t *testing.T) bool {
		if validatedBlk == nil {
			t.Errorf("the block should be validated after the second user")
			return false
		}

		retBlk := validatedBlk.GetBlock().GetBlock()
		retLeaderSigs := validatedBlk.GetSignatures().GetSignatures()

		if !reflect.DeepEqual(retBlk, blk) {
			t.Errorf("the block inside the validated block is invalid")
			return false
		}

		if len(retLeaderSigs) != 2 {
			t.Errorf("there must be 2 leader user signatures in the validated block.  Returned amount: %d", len(retLeaderSigs))
			return false
		}

		retFirstUser := retLeaderSigs[0].GetUser()
		retSecondUser := retLeaderSigs[1].GetUser()

		if !reflect.DeepEqual(retFirstUser, firstUser) {
			t.Errorf("the user in the first leader signature is invalid")
			return false
		}

		if !reflect.DeepEqual(retSecondUser, secondUser) {
			t.Errorf("the user in the second leader signature is invalid")
			return false
		}

		return true
	}

	//variables:
	createdOn := time.Now().UTC()
	waitBeforeRemovalTs := time.Duration(time.Second * 20)

	//create a block:
	blk := concrete_block.CreateBlockForTests()

	//create users:
	invalidUser, invalidPK := generateUserAndPKFn()
	firstUser, firstPK := generateUserAndPKFn()
	secondUser, secondPK := generateUserAndPKFn()
	thirdUser, thirdPK := generateUserAndPKFn()

	firstUserStake := mathrand.Float64() + 5.0
	secondUserStake := mathrand.Float64() + 10.0
	thirdUserStake := mathrand.Float64() + 2.0
	usersStake := map[string]float64{
		firstUser.GetMetaData().GetID().String():  firstUserStake,
		secondUser.GetMetaData().GetID().String(): secondUserStake,
		thirdUser.GetMetaData().GetID().String():  thirdUserStake,
	}

	neededStake := firstUserStake + (secondUserStake / 2)

	//create factories:
	userSigBuilderFactory := concrete_users.CreateSignatureBuilderFactoryForTests()
	userSigsBuilderFactory := concrete_users.CreateSignaturesBuilderFactoryForTests()
	signedBlockBuilderFactory := concrete_block.CreateSignedBlockBuilderFactoryForTests()
	valBlkBuilderFactory := concrete_validated_block.CreateBlockBuilderFactoryForTests()

	//sign the blocks:
	invalidSigID := uuid.NewV4()
	invalidSigCrOn := time.Now().UTC()
	invalidSig, invalidSigErr := userSigBuilderFactory.Create().Create().WithID(&invalidSigID).CreatedOn(invalidSigCrOn).WithUser(invalidUser).WithPrivateKey(invalidPK).WithInterface(blk).Now()
	if invalidSigErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", invalidSigErr.Error())
		return
	}

	invalidSignedBlkID := uuid.NewV4()
	invalidSignedBlk, invalidSignedBlkErr := signedBlockBuilderFactory.Create().Create().WithID(&invalidSignedBlkID).CreatedOn(createdOn).WithBlock(blk).WithSignature(invalidSig).Now()
	if invalidSignedBlkErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", invalidSignedBlkErr.Error())
		return
	}

	firstSigID := uuid.NewV4()
	firstSigCrOn := time.Now().UTC()
	firstSig, firstSigErr := userSigBuilderFactory.Create().Create().WithID(&firstSigID).CreatedOn(firstSigCrOn).WithUser(firstUser).WithPrivateKey(firstPK).WithInterface(blk).Now()
	if firstSigErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", firstSigErr.Error())
		return
	}

	firstSignedBlkID := uuid.NewV4()
	firstSignedBlk, firstSignedBlkErr := signedBlockBuilderFactory.Create().Create().WithID(&firstSignedBlkID).CreatedOn(createdOn).WithBlock(blk).WithSignature(firstSig).Now()
	if firstSignedBlkErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", firstSignedBlkErr.Error())
		return
	}

	secondSigID := uuid.NewV4()
	secondSigCrOn := time.Now().UTC()
	secondSig, secondSigErr := userSigBuilderFactory.Create().Create().WithID(&secondSigID).CreatedOn(secondSigCrOn).WithUser(secondUser).WithPrivateKey(secondPK).WithInterface(blk).Now()
	if secondSigErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", secondSigErr.Error())
		return
	}

	secondSignedBlkID := uuid.NewV4()
	secondSignedBlk, secondSignedBlkErr := signedBlockBuilderFactory.Create().Create().WithID(&secondSignedBlkID).CreatedOn(createdOn).WithBlock(blk).WithSignature(secondSig).Now()
	if secondSignedBlkErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", secondSignedBlkErr.Error())
		return
	}

	thirdSigID := uuid.NewV4()
	thirdSigCrOn := time.Now().UTC()
	thirdSig, thirdSigErr := userSigBuilderFactory.Create().Create().WithID(&thirdSigID).CreatedOn(thirdSigCrOn).WithUser(thirdUser).WithPrivateKey(thirdPK).WithInterface(blk).Now()
	if thirdSigErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", thirdSigErr.Error())
		return
	}

	thirdSignedBlkID := uuid.NewV4()
	thirdSignedBlk, thirdSignedBlkErr := signedBlockBuilderFactory.Create().Create().WithID(&thirdSignedBlkID).CreatedOn(createdOn).WithBlock(blk).WithSignature(thirdSig).Now()
	if thirdSignedBlkErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", thirdSignedBlkErr.Error())
		return
	}

	//create the channels:
	newSignedBlks := make(chan blocks.SignedBlock, 20)
	newValidatedBlks := make(chan validated_blocks.Block, 20)

	//create the application:
	validateApp := CreateValidateBlock(valBlkBuilderFactory, userSigsBuilderFactory, neededStake, usersStake, waitBeforeRemovalTs, newSignedBlks, newValidatedBlks)
	defer validateApp.Stop()

	//execute:
	go validateApp.Execute()

	//no signed block have been sent yet:
	preValidatedBlk := getValidatedBlockFn(newValidatedBlks)
	if preValidatedBlk != nil {
		t.Errorf("there should be no validated block since no signed block were submitted")
		return
	}

	//add the invalid signed block (the user is not in the stake list):
	newSignedBlks <- invalidSignedBlk
	time.Sleep(time.Second)
	invalidValidatedBlk := getValidatedBlockFn(newValidatedBlks)
	if invalidValidatedBlk != nil {
		t.Errorf("there should be no validated block because the signed block was created by a user NOT in the stake list")
		return
	}

	//add the signed blocks to the signed block channel:
	newSignedBlks <- firstSignedBlk
	time.Sleep(time.Second)
	firstValidatedBlk := getValidatedBlockFn(newValidatedBlks)
	if firstValidatedBlk != nil {
		t.Errorf("the validated block should not have been validated after the first user submission")
		return
	}

	newSignedBlks <- secondSignedBlk
	time.Sleep(time.Second)
	secondValidatedBlk := getValidatedBlockFn(newValidatedBlks)
	if !verifyValidatedBlockFn(secondValidatedBlk, blk, firstUser, secondUser, t) {
		return
	}

	newSignedBlks <- thirdSignedBlk
	time.Sleep(time.Second)
	thirdValidatedBlk := getValidatedBlockFn(newValidatedBlks)
	if thirdValidatedBlk != nil {
		t.Errorf("no new validated block should be pushed after the third submission")
		return
	}

	//wait enough to execute cleanup in the app:
	time.Sleep(waitBeforeRemovalTs * 2)

	//re-submit the older signed block.  This should regenerate another validated block:
	newSignedBlks <- firstSignedBlk
	time.Sleep(time.Second)
	againFirstValidatedBlk := getValidatedBlockFn(newValidatedBlks)
	if againFirstValidatedBlk != nil {
		t.Errorf("the validated block should not have been validated after the first user re-submission")
		return
	}

	newSignedBlks <- secondSignedBlk
	time.Sleep(time.Second)
	againValidatedBlk := getValidatedBlockFn(newValidatedBlks)
	if !verifyValidatedBlockFn(againValidatedBlk, blk, firstUser, secondUser, t) {
		return
	}

	newSignedBlks <- thirdSignedBlk
	time.Sleep(time.Second)
	lastThirdValidatedBlk := getValidatedBlockFn(newValidatedBlks)
	if lastThirdValidatedBlk != nil {
		t.Errorf("no new validated block should be pushed after the third submission")
		return
	}

}
