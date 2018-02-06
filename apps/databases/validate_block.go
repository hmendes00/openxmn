package databases

import (
	"log"
	"time"

	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	validated_blocks "github.com/XMNBlockchain/core/packages/lives/blocks/validated/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
)

// ValidateBlock represents an application that received SignedBlocks and created validated Blocks
type ValidateBlock struct {
	valBlkBuilderFactory validated_blocks.BlockBuilderFactory
	usersKarma           map[string]int
	toRemoveSignedBlks   map[string]time.Time
	signedBlks           map[string][]blocks.SignedBlock
	newSignedBlks        <-chan blocks.SignedBlock
	newValidatedBlks     chan<- validated_blocks.Block
	stop                 bool
	waitBeforeRemovalTs  time.Duration
}

// CreateValidateBlock creates a new ValidateBlock instance
func CreateValidateBlock(
	valBlkBuilderFactory validated_blocks.BlockBuilderFactory,
	usersKarma map[string]int,
	waitBeforeRemovalTs time.Duration,
	newSignedBlks <-chan blocks.SignedBlock,
	newValidatedBlks chan<- validated_blocks.Block,
) *ValidateBlock {
	out := ValidateBlock{
		valBlkBuilderFactory: valBlkBuilderFactory,
		usersKarma:           usersKarma,
		waitBeforeRemovalTs:  waitBeforeRemovalTs,
		toRemoveSignedBlks:   map[string]time.Time{},
		signedBlks:           map[string][]blocks.SignedBlock{},
		newSignedBlks:        newSignedBlks,
		newValidatedBlks:     newValidatedBlks,
		stop:                 false,
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
			hashAsString := oneSignedBlk.GetBlock().GetHashTree().GetHash().String()

			//if the has exists in the toRemove map:
			if _, ok := blk.toRemoveSignedBlks[hashAsString]; ok {
				log.Printf("the block (hash: %s) already got processed successfully.  Skipping.\n", hashAsString)
				break
			}

			//if the hash already exists in the map:
			if _, ok := blk.signedBlks[hashAsString]; ok {
				//writing the logs:
				userID := oneSignedBlk.GetSignature().GetUser().GetID().String()
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

		//verify if signed blocks got validated by enough karma:
		for signedBlkHash, signedBlks := range blk.signedBlks {
			//get the amount of karma:
			receivedKarma := 0
			leaderSignatures := []users.Signature{}
			for _, oneSignedBlk := range signedBlks {

				//retrieve the leader signature:
				leaderSig := oneSignedBlk.GetSignature()

				//add the received karma:
				userIDAsString := leaderSig.GetUser().GetID().String()
				if userKarma, ok := blk.usersKarma[userIDAsString]; ok {
					//add the karma:
					receivedKarma += userKarma

					//add the leader signature in the list, since the karma was verified:
					leaderSignatures = append(leaderSignatures, leaderSig)

					continue
				}
			}

			//if we have enough karma, build the validated block and remove the signedBlks from the list:
			firstBlk := signedBlks[0].GetBlock()
			neededKarma := firstBlk.GetNeededKarma()
			if receivedKarma >= neededKarma {
				validatedBlk, validatedBlkErr := blk.valBlkBuilderFactory.Create().Create().WithBlock(signedBlks[0]).WithSignatures(leaderSignatures).Now()
				if validatedBlkErr != nil {
					log.Printf("there was an error while building a validated block: %s", validatedBlkErr.Error())
					continue
				}

				//write to the logs:
				log.Printf("the block (hash: %s) needed %d karma to be verified, received: %d\n", signedBlkHash, neededKarma, receivedKarma)

				//add the validated block to the channel:
				blk.newValidatedBlks <- validatedBlk

				//add the hash to the list of element to remove:
				blk.toRemoveSignedBlks[signedBlkHash] = time.Now()
			}

		}

	}

}
