package leaders

import (
	"log"
	"time"

	blocks "github.com/XMNBlockchain/core/packages/blockchains/blocks/blocks/domain"
	aggr_trs "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
	uuid "github.com/satori/go.uuid"
)

// Push represents an application that take the tranactions received from the API, to the leader
type Push struct {
	blkBuilderFactory blocks.BlockBuilderFactory
	blkTimeDuration   time.Duration
	newSignedAggrTrs  <-chan aggr_trs.SignedTransactions
	newBlk            chan<- blocks.Block
	stop              bool
}

// CreatePush creates a new Push instance
func CreatePush(
	blkBuilderFactory blocks.BlockBuilderFactory,
	blkTimeDuration time.Duration,
	newSignedAggrTrs <-chan aggr_trs.SignedTransactions,
	newBlk chan<- blocks.Block,
	stop bool,
) *Push {
	out := Push{
		blkBuilderFactory: blkBuilderFactory,
		blkTimeDuration:   blkTimeDuration,
		newSignedAggrTrs:  newSignedAggrTrs,
		newBlk:            newBlk,
		stop:              false,
	}

	return &out
}

// Stop stops the push signed aggregated transactions application
func (pu *Push) Stop() {
	pu.stop = true
}

// Execute executes the push signed aggregated transactions application
func (pu *Push) Execute() {

	curTime := time.Now().UTC()
	aggrSignedTrs := []aggr_trs.SignedTransactions{}

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
		aggrSignedTrs = []aggr_trs.SignedTransactions{}
	}

}
