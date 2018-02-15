package databases

import (
	"crypto/rand"
	"crypto/rsa"
	"reflect"
	"testing"
	"time"

	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	concrete_block "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/infrastructure"
	validated_blocks "github.com/XMNBlockchain/core/packages/lives/blocks/validated/domain"
	concrete_validated_block "github.com/XMNBlockchain/core/packages/lives/blocks/validated/infrastructure"
	concrete_hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/users/infrastructure"
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
		retLeaderSigs := validatedBlk.GetLeaderSignatures()

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
	waitBeforeRemovalTs := time.Duration(time.Second * 20)

	//create a block:
	blk := concrete_block.CreateBlockForTests(t)

	//user IDs and karma:
	invalidUserID := uuid.NewV4()
	firstUserID := uuid.NewV4()
	secondUserID := uuid.NewV4()
	thirdUserID := uuid.NewV4()
	usersKarma := map[string]int{
		firstUserID.String():  blk.GetNeededKarma() / 2,
		secondUserID.String(): blk.GetNeededKarma() / 2,
		thirdUserID.String():  1,
	}

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
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	valBlkBuilderFactory := concrete_validated_block.CreateBlockBuilderFactory(htBuilderFactory)

	//sign the blocks:
	invalidSig, _ := userSigBuilderFactory.Create().Create().WithUserID(&invalidUserID).WithPrivateKey(invalidPK).WithInterface(blk).Now()
	invalidSignedBlk, _ := signedBlockBuilderFactory.Create().Create().WithBlock(blk).WithSignature(invalidSig).Now()
	firstSig, _ := userSigBuilderFactory.Create().Create().WithUserID(&firstUserID).WithPrivateKey(firstPK).WithInterface(blk).Now()
	firstSignedBlk, _ := signedBlockBuilderFactory.Create().Create().WithBlock(blk).WithSignature(firstSig).Now()
	secondSig, _ := userSigBuilderFactory.Create().Create().WithUserID(&secondUserID).WithPrivateKey(secondPK).WithInterface(blk).Now()
	secondSignedBlk, _ := signedBlockBuilderFactory.Create().Create().WithBlock(blk).WithSignature(secondSig).Now()
	thirdSig, _ := userSigBuilderFactory.Create().Create().WithUserID(&thirdUserID).WithPrivateKey(thirdPK).WithInterface(blk).Now()
	thirdSignedBlk, _ := signedBlockBuilderFactory.Create().Create().WithBlock(blk).WithSignature(thirdSig).Now()

	//create the channels:
	newSignedBlks := make(chan blocks.SignedBlock, 20)
	newValidatedBlks := make(chan validated_blocks.Block, 20)

	//create the application:
	validateApp := CreateValidateBlock(valBlkBuilderFactory, usersKarma, waitBeforeRemovalTs, newSignedBlks, newValidatedBlks)
	defer validateApp.Stop()

	//execute:
	go validateApp.Execute()

	//no signed block have been sent yet:
	preValidatedBlk := getValidatedBlockFn(newValidatedBlks)
	if preValidatedBlk != nil {
		t.Errorf("there should be no validated block since no signed block were submitted")
		return
	}

	//add the invalid signed block (the user is not in the karma list:
	newSignedBlks <- invalidSignedBlk
	time.Sleep(time.Second)
	invalidValidatedBlk := getValidatedBlockFn(newValidatedBlks)
	if invalidValidatedBlk != nil {
		t.Errorf("there should be no validated block because the signed block was created by a user NOT in the karma list")
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
