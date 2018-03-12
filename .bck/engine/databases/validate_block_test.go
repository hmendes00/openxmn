package databases

import (
	"crypto/rand"
	"crypto/rsa"
	mathrand "math/rand"
	"reflect"
	"testing"
	"time"

	blocks "github.com/XMNBlockchain/core/packages/blockchains/blocks/blocks/domain"
	concrete_block "github.com/XMNBlockchain/core/packages/blockchains/blocks/blocks/infrastructure"
	validated_blocks "github.com/XMNBlockchain/core/packages/blockchains/blocks/validated/domain"
	concrete_validated_block "github.com/XMNBlockchain/core/packages/blockchains/blocks/validated/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/blockchains/users/infrastructure"
	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	uuid "github.com/satori/go.uuid"
)

func TestValidateBlock_Success(t *testing.T) {

	t.Parallel()

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
	verifyValidatedBlockFn := func(validatedBlk validated_blocks.Block, blk blocks.Block, firstUserID *uuid.UUID, secondUserID *uuid.UUID, t *testing.T) bool {
		if validatedBlk == nil {
			t.Errorf("the block should be validated after the second user")
			return false
		}

		retBlk := validatedBlk.GetBlock().GetBlock()
		retLeaderSigs := validatedBlk.GetSignatures()

		if !reflect.DeepEqual(retBlk, blk) {
			t.Errorf("the block inside the validated block is invalid")
			return false
		}

		if len(retLeaderSigs) != 2 {
			t.Errorf("there must be 2 leader user signatures in the validated block.  Returned amount: %d", len(retLeaderSigs))
			return false
		}

		retFirstUserID := retLeaderSigs[0].GetUser().GetID()
		retSecondUserID := retLeaderSigs[1].GetUser().GetID()

		if !reflect.DeepEqual(retFirstUserID, firstUserID) {
			t.Errorf("the userID in the first leader signature is invalid.  Expected: %s, Returned: %s", firstUserID.String(), retFirstUserID.String())
			return false
		}

		if !reflect.DeepEqual(retSecondUserID, secondUserID) {
			t.Errorf("the userID in the first leader signature is invalid.  Expected: %s, Returned: %s", secondUserID.String(), retSecondUserID.String())
			return false
		}

		return true
	}

	//variables:
	createdOn := time.Now().UTC()
	waitBeforeRemovalTs := time.Duration(time.Second * 20)

	//create a block:
	blk := concrete_block.CreateBlockForTests(t)

	//user IDs and stake:
	invalidUserID := uuid.NewV4()
	firstUserID := uuid.NewV4()
	secondUserID := uuid.NewV4()
	thirdUserID := uuid.NewV4()

	firstUserStake := mathrand.Float64() + 5.0
	secondUserStake := mathrand.Float64() + 10.0
	thirdUserStake := mathrand.Float64() + 2.0
	usersStake := map[string]float64{
		firstUserID.String():  firstUserStake,
		secondUserID.String(): secondUserStake,
		thirdUserID.String():  thirdUserStake,
	}

	neededStake := firstUserStake + (secondUserStake / 2)

	//create the private keys:
	reader := rand.Reader
	bitSize := 4096
	invalidRawPK, _ := rsa.GenerateKey(reader, bitSize)
	invalidPK, _ := concrete_cryptography.CreatePrivateKeyBuilderFactory().Create().WithKey(invalidRawPK).Now()
	firstRawPK, _ := rsa.GenerateKey(reader, bitSize)
	firstPK, _ := concrete_cryptography.CreatePrivateKeyBuilderFactory().Create().WithKey(firstRawPK).Now()
	secondRawPK, _ := rsa.GenerateKey(reader, bitSize)
	secondPK, _ := concrete_cryptography.CreatePrivateKeyBuilderFactory().Create().WithKey(secondRawPK).Now()
	thirdRawPK, _ := rsa.GenerateKey(reader, bitSize)
	thirdPK, _ := concrete_cryptography.CreatePrivateKeyBuilderFactory().Create().WithKey(thirdRawPK).Now()

	//create factories:
	publicKeyBuilderFactory := concrete_cryptography.CreatePublicKeyBuilderFactory()
	sigBuilderFactory := concrete_cryptography.CreateSignatureBuilderFactory(publicKeyBuilderFactory)
	userBuilderFactory := concrete_users.CreateUserBuilderFactory()
	userSigBuilderFactory := concrete_users.CreateSignatureBuilderFactory(sigBuilderFactory, userBuilderFactory)
	signedBlockBuilderFactory := concrete_block.CreateSignedBlockBuilderFactory()
	valBlkBuilderFactory := concrete_validated_block.CreateBlockBuilderFactory()

	//sign the blocks:
	invalidSig, invalidSigErr := userSigBuilderFactory.Create().Create().WithUserID(&invalidUserID).WithPrivateKey(invalidPK).WithInterface(blk).Now()
	if invalidSigErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", invalidSigErr.Error())
	}

	invalidSignedBlkID := uuid.NewV4()
	invalidSignedBlk, invalidSignedBlkErr := signedBlockBuilderFactory.Create().Create().WithID(&invalidSignedBlkID).CreatedOn(createdOn).WithBlock(blk).WithSignature(invalidSig).Now()
	if invalidSignedBlkErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", invalidSignedBlkErr.Error())
	}

	firstSig, firstSigErr := userSigBuilderFactory.Create().Create().WithUserID(&firstUserID).WithPrivateKey(firstPK).WithInterface(blk).Now()
	if firstSigErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", firstSigErr.Error())
	}

	firstSignedBlkID := uuid.NewV4()
	firstSignedBlk, firstSignedBlkErr := signedBlockBuilderFactory.Create().Create().WithID(&firstSignedBlkID).CreatedOn(createdOn).WithBlock(blk).WithSignature(firstSig).Now()
	if firstSignedBlkErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", firstSignedBlkErr.Error())
	}

	secondSig, secondSigErr := userSigBuilderFactory.Create().Create().WithUserID(&secondUserID).WithPrivateKey(secondPK).WithInterface(blk).Now()
	if secondSigErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", secondSigErr.Error())
	}

	secondSignedBlkID := uuid.NewV4()
	secondSignedBlk, secondSignedBlkErr := signedBlockBuilderFactory.Create().Create().WithID(&secondSignedBlkID).CreatedOn(createdOn).WithBlock(blk).WithSignature(secondSig).Now()
	if secondSignedBlkErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", secondSignedBlkErr.Error())
	}

	thirdSig, thirdSigErr := userSigBuilderFactory.Create().Create().WithUserID(&thirdUserID).WithPrivateKey(thirdPK).WithInterface(blk).Now()
	if thirdSigErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", thirdSigErr.Error())
	}

	thirdSignedBlkID := uuid.NewV4()
	thirdSignedBlk, thirdSignedBlkErr := signedBlockBuilderFactory.Create().Create().WithID(&thirdSignedBlkID).CreatedOn(createdOn).WithBlock(blk).WithSignature(thirdSig).Now()
	if thirdSignedBlkErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", thirdSignedBlkErr.Error())
	}

	//create the channels:
	newSignedBlks := make(chan blocks.SignedBlock, 20)
	newValidatedBlks := make(chan validated_blocks.Block, 20)

	//create the application:
	validateApp := CreateValidateBlock(valBlkBuilderFactory, neededStake, usersStake, waitBeforeRemovalTs, newSignedBlks, newValidatedBlks)
	defer validateApp.Stop()

	//execute:
	go validateApp.Execute()

	//no signed block have been sent yet:
	preValidatedBlk := getValidatedBlockFn(newValidatedBlks)
	if preValidatedBlk != nil {
		t.Errorf("there should be no validated block since no signed block were submitted")
		return
	}

	//add the invalid signed block (the user is not in the stake list:
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
	if !verifyValidatedBlockFn(secondValidatedBlk, blk, &firstUserID, &secondUserID, t) {
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
	if !verifyValidatedBlockFn(againValidatedBlk, blk, &firstUserID, &secondUserID, t) {
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
