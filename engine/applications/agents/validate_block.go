package agents

import (
	"log"
	"time"

	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks"
	validated_blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks/validated"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	uuid "github.com/satori/go.uuid"
)

// ValidateBlock represents an application that received SignedBlocks and created validated Blocks
type ValidateBlock struct {
	pk                         cryptography.PrivateKey
	user                       users.User
	userSigBuilderFactory      users.SignatureBuilderFactory
	valBlkBuilderFactory       validated_blocks.BlockBuilderFactory
	signedValBlkBuilderFactory validated_blocks.SignedBlockBuilderFactory
	userSigsBuilderFactory     users.SignaturesBuilderFactory
	neededStakerPerBlk         float64
	usersStake                 map[string]float64
	toRemoveSignedBlks         map[string]time.Time
	signedBlks                 map[string][]blocks.SignedBlock
	newSignedBlks              <-chan blocks.SignedBlock
	newSignedValidatedBlk      chan<- validated_blocks.SignedBlock
	stop                       bool
	waitBeforeRemovalTs        time.Duration
}

// CreateValidateBlock creates a new ValidateBlock instance
func CreateValidateBlock(
	pk cryptography.PrivateKey,
	user users.User,
	userSigBuilderFactory users.SignatureBuilderFactory,
	valBlkBuilderFactory validated_blocks.BlockBuilderFactory,
	signedValBlkBuilderFactory validated_blocks.SignedBlockBuilderFactory,
	userSigsBuilderFactory users.SignaturesBuilderFactory,
	neededStakerPerBlk float64,
	usersStake map[string]float64,
	waitBeforeRemovalTs time.Duration,
	newSignedBlks <-chan blocks.SignedBlock,
	newSignedValidatedBlk chan<- validated_blocks.SignedBlock,
) *ValidateBlock {
	out := ValidateBlock{
		pk:   pk,
		user: user,
		userSigBuilderFactory:      userSigBuilderFactory,
		valBlkBuilderFactory:       valBlkBuilderFactory,
		signedValBlkBuilderFactory: signedValBlkBuilderFactory,
		userSigsBuilderFactory:     userSigsBuilderFactory,
		neededStakerPerBlk:         neededStakerPerBlk,
		usersStake:                 usersStake,
		waitBeforeRemovalTs:        waitBeforeRemovalTs,
		toRemoveSignedBlks:         map[string]time.Time{},
		signedBlks:                 map[string][]blocks.SignedBlock{},
		newSignedBlks:              newSignedBlks,
		newSignedValidatedBlk:      newSignedValidatedBlk,
		stop: false,
	}

	return &out
}

// Stop stops the validate block application
func (blk *ValidateBlock) Stop() {
	blk.stop = true
}

// Execute executes the validate block application
func (blk *ValidateBlock) Execute() {

	for {

		//verify if the app is stopped:
		if blk.stop {
			log.Println("stopping...")
			return
		}

		//retrieve the new signed blocks:
		select {
		case oneSignedBlk := <-blk.newSignedBlks:

			//get the hash of the block:
			hashAsString := oneSignedBlk.GetBlock().GetMetaData().GetHashTree().GetHash().String()

			//if the hash exists in the toRemove map:
			if _, ok := blk.toRemoveSignedBlks[hashAsString]; ok {
				log.Printf("the block (hash: %s) already got processed successfully.  Skipping.\n", hashAsString)
				break
			}

			//if the hash already exists in the map:
			if _, ok := blk.signedBlks[hashAsString]; ok {
				//writing the logs:
				userID := oneSignedBlk.GetSignature().GetUser().GetMetaData().GetID().String()
				log.Printf("Adding a new verification from userID (%s) to the block (hash: %s)\n", userID, hashAsString)

				//adding the signed block to the list:
				blk.signedBlks[hashAsString] = append(blk.signedBlks[hashAsString], oneSignedBlk)
				break
			}

			//writing the logs:
			log.Printf("Adding a block (hash: %s) to the list of block to verify.\n", hashAsString)

			//the hash is not in our map yet, so add it:
			blk.signedBlks[hashAsString] = []blocks.SignedBlock{
				oneSignedBlk,
			}

			break
		}

		//if there is signed blocks to remove:
		if len(blk.toRemoveSignedBlks) > 0 {
			currentTs := time.Now()
			for oneHash, oneTs := range blk.toRemoveSignedBlks {
				//if we waited enough, delete the signed block:
				if oneTs.Add(blk.waitBeforeRemovalTs).After(currentTs) {
					//log:
					log.Printf("Removing the block from the list (hash: %s)", oneHash)

					//delete:
					delete(blk.signedBlks, oneHash)
					delete(blk.toRemoveSignedBlks, oneHash)
				}
			}
		}

		//if there is no signed blocks, continue:
		if len(blk.signedBlks) <= 0 {
			continue
		}

		//verify if signed blocks got validated by enough stake:
		for signedBlkHash, signedBlks := range blk.signedBlks {
			//get the amount of stake:
			receivedStake := 0.0
			leaderSignatures := []users.Signature{}
			for _, oneSignedBlk := range signedBlks {

				//retrieve the leader signature:
				leaderSig := oneSignedBlk.GetSignature()

				//add the received stake:
				userIDAsString := leaderSig.GetUser().GetMetaData().GetID().String()
				if userKarma, ok := blk.usersStake[userIDAsString]; ok {
					//add the stake:
					receivedStake += userKarma

					//add the leader signature in the list, since the stake was verified:
					leaderSignatures = append(leaderSignatures, leaderSig)

					continue
				}
			}

			//if we have enough stake, build the validated block and remove the signedBlks from the list:
			if receivedStake >= blk.neededStakerPerBlk {

				userSigsID := uuid.NewV4()
				userSigsCrOn := time.Now().UTC()
				userSigs, userSigsErr := blk.userSigsBuilderFactory.Create().Create().WithID(&userSigsID).CreatedOn(userSigsCrOn).WithSignatures(leaderSignatures).Now()
				if userSigsErr != nil {
					log.Printf("the user signatures instance could not be built: %s", userSigsErr.Error())
					continue
				}

				validatedBlkID := uuid.NewV4()
				validatedBlkTs := time.Now().UTC()
				validatedBlk, validatedBlkErr := blk.valBlkBuilderFactory.Create().Create().WithID(&validatedBlkID).CreatedOn(validatedBlkTs).WithBlock(signedBlks[0]).WithSignatures(userSigs).Now()
				if validatedBlkErr != nil {
					log.Printf("there was an error while building a validated block: %s", validatedBlkErr.Error())
					continue
				}

				//write to the logs:
				log.Printf("the block (hash: %s) needed %f stake to be verified, received: %f\n", signedBlkHash, blk.neededStakerPerBlk, receivedStake)

				//sign the validated block:
				blkSigID := uuid.NewV4()
				blkSigCrOn := time.Now().UTC()
				blkSig, blkSigErr := blk.userSigBuilderFactory.Create().Create().WithID(&blkSigID).CreatedOn(blkSigCrOn).WithUser(blk.user).WithPrivateKey(blk.pk).WithInterface(validatedBlk).Now()
				if blkSigErr != nil {
					log.Printf("there was an error while signing the validated block (ID: %s): %s", validatedBlk.GetMetaData().GetID().String(), blkSigErr.Error())
					continue
				}

				signedValBlkID := uuid.NewV4()
				signedValBlkCrOn := time.Now().UTC()
				signedValBlk, signedValBlkErr := blk.signedValBlkBuilderFactory.Create().Create().WithID(&signedValBlkID).CreatedOn(signedValBlkCrOn).WithSignature(blkSig).WithBlock(validatedBlk).Now()
				if signedValBlkErr != nil {
					log.Printf("there was an error while building a signed validated block: %s", signedValBlkErr.Error())
					continue
				}

				//add the validated block to the channel:
				blk.newSignedValidatedBlk <- signedValBlk

				//add the hash to the list of element to remove:
				blk.toRemoveSignedBlks[signedBlkHash] = time.Now()
			}

		}

	}

}
