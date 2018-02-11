package infrastructure

import (
	aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/domain"
	signed "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
)

// TransactionsRepository represents a concrete TransactionsRepository implementation
type TransactionsRepository struct {
	signedTrsRepository         signed.TransactionRepository
	atomicSignedTrsRepository   signed.AtomicTransactionRepository
	aggregatedTrsBuilderFactory aggregated.TransactionsBuilderFactory
}

// CreateTransactionsRepository creates a new TransactionsRepository instance
func CreateTransactionsRepository(signedTrsRepository signed.TransactionRepository, atomicSignedTrsRepository signed.AtomicTransactionRepository, aggregatedTrsBuilderFactory aggregated.TransactionsBuilderFactory) aggregated.TransactionsRepository {
	out := TransactionsRepository{
		signedTrsRepository:         signedTrsRepository,
		atomicSignedTrsRepository:   atomicSignedTrsRepository,
		aggregatedTrsBuilderFactory: aggregatedTrsBuilderFactory,
	}
	return &out
}

// Retrieve retrieves an aggregated Transactions instance
func (rep *TransactionsRepository) Retrieve(dirPath string) (aggregated.Transactions, error) {

	//create the paths:
	/*trsPath := filepath.Join(dirPath, "signed_transactions")
		atomicTrsPath := filepath.Join(dirPath, "signed_atomic_transactions")

		//retrieve the signed transactions:
		signedTrs, signedTrsErr := rep.signedTrsRepository.RetrieveAll(trsPath)
		if signedTrsErr != nil {
			return nil, signedTrsErr
		}

		//retrieve the atomic signed transactions:
		atomicSignedTrs, atomicSignedTrsErr := rep.atomicSignedTrsRepository.RetrieveAll(atomicTrsPath)
		if atomicSignedTrsErr != nil {
			return nil, atomicSignedTrsErr
		}



	    //build the aggregated transactions:
	    aggrTrs, aggrTrsErr := rep.aggregatedTrsBuilderFactory.Create().Create().WithID(id).CreatedOn(ts)*/

	return nil, nil
}

// RetrieveAll retrieves an aggregated []Transactions instance
func (rep *TransactionsRepository) RetrieveAll(dirPath string) ([]aggregated.Transactions, error) {
	return nil, nil
}
