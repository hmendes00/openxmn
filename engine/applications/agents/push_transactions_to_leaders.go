package agents

import (
	"log"
	"time"

	signed_transactions "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/transactions/signed"
	aggregated_transactions "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/transactions/signed/aggregated"
	uuid "github.com/satori/go.uuid"
)

// PushTransactionsToLeanders represents an application that push the tranactions received from the API, to the leader
type PushTransactionsToLeanders struct {
	aggrTrsBuilderFactory           aggregated_transactions.TransactionsBuilderFactory
	signedTransBuilderFactory       signed_transactions.TransactionsBuilderFactory
	signedAtomicTransBuilderFactory signed_transactions.AtomicTransactionsBuilderFactory
	aggrTimeDuration                time.Duration
	newSignedTrs                    <-chan signed_transactions.Transaction
	newAtomicSignedTrs              <-chan signed_transactions.AtomicTransaction
	newAggregatedTrs                chan<- aggregated_transactions.Transactions
	stop                            bool
}

// CreatePushTransactionsToLeanders creates a new PushTransactionsToLeanders instance
func CreatePushTransactionsToLeanders(
	aggrTrsBuilderFactory aggregated_transactions.TransactionsBuilderFactory,
	signedTransBuilderFactory signed_transactions.TransactionsBuilderFactory,
	signedAtomicTransBuilderFactory signed_transactions.AtomicTransactionsBuilderFactory,
	aggrTimeDuration time.Duration,
	newSignedTrs <-chan signed_transactions.Transaction,
	newAtomicSignedTrs <-chan signed_transactions.AtomicTransaction,
	newAggregatedTrs chan<- aggregated_transactions.Transactions,
) *PushTransactionsToLeanders {
	out := PushTransactionsToLeanders{
		aggrTrsBuilderFactory:           aggrTrsBuilderFactory,
		signedTransBuilderFactory:       signedTransBuilderFactory,
		signedAtomicTransBuilderFactory: signedAtomicTransBuilderFactory,
		aggrTimeDuration:                aggrTimeDuration,
		newSignedTrs:                    newSignedTrs,
		newAtomicSignedTrs:              newAtomicSignedTrs,
		newAggregatedTrs:                newAggregatedTrs,
		stop:                            false,
	}

	return &out
}

// Stop stops the application
func (pu *PushTransactionsToLeanders) Stop() {
	pu.stop = true
}

// Execute executes the application
func (pu *PushTransactionsToLeanders) Execute() {

	curTime := time.Now().UTC()
	signedTrs := []signed_transactions.Transaction{}
	atomicSignedTrs := []signed_transactions.AtomicTransaction{}

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
			transID := uuid.NewV4()
			transCrOn := time.Now().UTC()
			signedTrans, signedTransErr := pu.signedTransBuilderFactory.Create().Create().WithID(&transID).CreatedOn(transCrOn).WithTransactions(signedTrs).Now()
			if signedTransErr != nil {
				log.Printf("the signed transactions instance could not be built: %s", signedTransErr.Error())
				continue
			}

			aggrTrsBuilder.WithTransactions(signedTrans)
		}

		if amountAtomicTrs > 0 {
			atomicTransID := uuid.NewV4()
			atomicTransCrOn := time.Now().UTC()
			atomicTrans, atomicTransErr := pu.signedAtomicTransBuilderFactory.Create().Create().WithID(&atomicTransID).CreatedOn(atomicTransCrOn).WithTransactions(atomicSignedTrs).Now()
			if atomicTransErr != nil {
				log.Printf("the signed atomic transactions instance could not be built: %s", atomicTransErr.Error())
				continue
			}

			aggrTrsBuilder.WithAtomicTransactions(atomicTrans)
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
		signedTrs = []signed_transactions.Transaction{}
		atomicSignedTrs = []signed_transactions.AtomicTransaction{}
	}

}
