package transactions

import (
	"log"
	"time"

	aggr_trs "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
	signed_trs "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/domain"
	uuid "github.com/satori/go.uuid"
)

// Push represents an application that take the tranactions received from the API, to the leader
type Push struct {
	aggrTrsBuilderFactory aggr_trs.TransactionsBuilderFactory
	aggrTimeDuration      time.Duration
	newSignedTrs          <-chan signed_trs.Transaction
	newAtomicSignedTrs    <-chan signed_trs.AtomicTransaction
	newAggregatedTrs      chan<- aggr_trs.Transactions
	stop                  bool
}

// CreatePush creates a new Push instance
func CreatePush(
	aggrTrsBuilderFactory aggr_trs.TransactionsBuilderFactory,
	aggrTimeDuration time.Duration,
	newSignedTrs <-chan signed_trs.Transaction,
	newAtomicSignedTrs <-chan signed_trs.AtomicTransaction,
	newAggregatedTrs chan<- aggr_trs.Transactions,
) *Push {
	out := Push{
		aggrTrsBuilderFactory: aggrTrsBuilderFactory,
		aggrTimeDuration:      aggrTimeDuration,
		newSignedTrs:          newSignedTrs,
		newAtomicSignedTrs:    newAtomicSignedTrs,
		newAggregatedTrs:      newAggregatedTrs,
		stop:                  false,
	}

	return &out
}

// Stop stops the push transactions application
func (pu *Push) Stop() {
	pu.stop = true
}

// Execute executes the push transactions application
func (pu *Push) Execute() {

	curTime := time.Now().UTC()
	signedTrs := []signed_trs.Transaction{}
	atomicSignedTrs := []signed_trs.AtomicTransaction{}

	for {

		//verify if the app is stopped:
		if pu.stop {
			log.Println("stopping...")
			return
		}

		select {
		case oneSignedTrs := <-pu.newSignedTrs:
			signedTrs = append(signedTrs, oneSignedTrs)
			break
		case oneSignedAtomicTrs := <-pu.newAtomicSignedTrs:
			atomicSignedTrs = append(atomicSignedTrs, oneSignedAtomicTrs)
			break
		}

		//if its time to aggregate the transactions:
		newTime := time.Now().UTC()
		if curTime.Add(pu.aggrTimeDuration).Before(newTime) {
			continue
		}

		//if there is no transactions, continue:
		amountAtomicTrs := len(atomicSignedTrs)
		amountTrs := len(signedTrs)
		if amountAtomicTrs <= 0 && amountTrs <= 0 {
			continue
		}

		//aggregate transactions:
		id := uuid.NewV4()
		ts := time.Now().UTC()
		aggrTrsBuilder := pu.aggrTrsBuilderFactory.Create().Create().WithID(&id).CreatedOn(ts)
		if amountTrs > 0 {
			aggrTrsBuilder.WithTransactions(signedTrs)
		}

		if amountAtomicTrs > 0 {
			aggrTrsBuilder.WithAtomicTransactions(atomicSignedTrs)
		}

		aggregatedTrs, aggregatedTrsErr := aggrTrsBuilder.Now()
		if aggregatedTrsErr != nil {
			log.Printf("the transactions could not be aggregated: %s", aggregatedTrsErr.Error())
			continue
		}

		//add the aggregated transaction to the channel:
		pu.newAggregatedTrs <- aggregatedTrs

		//reset:
		curTime = time.Now().UTC()
		signedTrs = []signed_trs.Transaction{}
		atomicSignedTrs = []signed_trs.AtomicTransaction{}
	}

}
