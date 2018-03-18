package agents

import (
	"log"
	"time"

	blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/blocks"
	aggregated_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/transactions/signed/aggregated"
	uuid "github.com/satori/go.uuid"
)

// PushAggregatedTransactionsFromLeadersToBlocks represents an application that takes the aggregated transactions from the leaders and push it to the blocks
type PushAggregatedTransactionsFromLeadersToBlocks struct {
	blkBuilderFactory blocks.BlockBuilderFactory
	blkTimeDuration   time.Duration
	newSignedAggrTrs  <-chan aggregated_transactions.SignedTransactions
	newBlk            chan<- blocks.Block
	stop              bool
}

// CreatePushAggregatedTransactionsFromLeadersToBlocks creates a new PushAggregatedTransactionsFromLeadersToBlocks instance
func CreatePushAggregatedTransactionsFromLeadersToBlocks(
	blkBuilderFactory blocks.BlockBuilderFactory,
	blkTimeDuration time.Duration,
	newSignedAggrTrs <-chan aggregated_transactions.SignedTransactions,
	newBlk chan<- blocks.Block,
	stop bool,
) *PushAggregatedTransactionsFromLeadersToBlocks {
	out := PushAggregatedTransactionsFromLeadersToBlocks{
		blkBuilderFactory: blkBuilderFactory,
		blkTimeDuration:   blkTimeDuration,
		newSignedAggrTrs:  newSignedAggrTrs,
		newBlk:            newBlk,
		stop:              false,
	}

	return &out
}

// Stop stops the application
func (pu *PushAggregatedTransactionsFromLeadersToBlocks) Stop() {
	pu.stop = true
}

// Execute executes the application
func (pu *PushAggregatedTransactionsFromLeadersToBlocks) Execute() {

	curTime := time.Now().UTC()
	aggrSignedTrs := []aggregated_transactions.SignedTransactions{}

	for {

		//verify if the app is stopped:
		if pu.stop {
			log.Println("stopping...")
			return
		}

		select {
		case oneSignedAggrTrs := <-pu.newSignedAggrTrs:
			aggrSignedTrs = append(aggrSignedTrs, oneSignedAggrTrs)
			break
		}

		//if its time to create the block:
		newTime := time.Now().UTC()
		if curTime.Add(pu.blkTimeDuration).Before(newTime) {
			continue
		}

		//if there is no transactions, continue:
		if len(aggrSignedTrs) <= 0 {
			continue
		}

		//create the block:
		id := uuid.NewV4()
		ts := time.Now().UTC()
		blk, blkErr := pu.blkBuilderFactory.Create().Create().WithID(&id).CreatedOn(ts).WithTransactions(aggrSignedTrs).Now()
		if blkErr != nil {
			log.Printf("the block could not be created: %s", blkErr.Error())
			continue
		}

		//add the block to the channel:
		pu.newBlk <- blk

		//reset:
		curTime = time.Now().UTC()
		aggrSignedTrs = []aggregated_transactions.SignedTransactions{}
	}

}
