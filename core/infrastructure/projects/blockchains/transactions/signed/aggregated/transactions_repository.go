package aggregated

import (
	"path/filepath"

	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
	signed "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/transactions/signed"
	aggregated "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/transactions/signed/aggregated"
)

// TransactionsRepository represents a concrete TransactionsRepository implementation
type TransactionsRepository struct {
	metaDataRepository          metadata.MetaDataRepository
	signedTrsRepository         signed.TransactionsRepository
	signedAtomicTrsRepository   signed.AtomicTransactionsRepository
	aggregatedTrsBuilderFactory aggregated.TransactionsBuilderFactory
}

// CreateTransactionsRepository creates a new TransactionsRepository instance
func CreateTransactionsRepository(
	metaDataRepository metadata.MetaDataRepository,
	signedTrsRepository signed.TransactionsRepository,
	signedAtomicTrsRepository signed.AtomicTransactionsRepository,
	aggregatedTrsBuilderFactory aggregated.TransactionsBuilderFactory,
) aggregated.TransactionsRepository {
	out := TransactionsRepository{
		metaDataRepository:          metaDataRepository,
		signedTrsRepository:         signedTrsRepository,
		signedAtomicTrsRepository:   signedAtomicTrsRepository,
		aggregatedTrsBuilderFactory: aggregatedTrsBuilderFactory,
	}
	return &out
}

// Retrieve retrieves an aggregated Transactions instance
func (rep *TransactionsRepository) Retrieve(dirPath string) (aggregated.Transactions, error) {
	//retrieve the metadata:
	met, metErr := rep.metaDataRepository.Retrieve(dirPath)
	if metErr != nil {
		return nil, metErr
	}

	//create the aggregated transactions builder:
	aggrTrsBuilder := rep.aggregatedTrsBuilderFactory.Create().WithMetaData(met)

	//retrieve the transactions, if any:
	trsPath := filepath.Join(dirPath, "signed_transactions")
	trs, trsErr := rep.signedTrsRepository.Retrieve(trsPath)
	if trsErr == nil {
		aggrTrsBuilder.WithTransactions(trs)
	}

	//retrieve the atomic transactions, if any:
	atomicTrsPath := filepath.Join(dirPath, "signed_atomic_transactions")
	atomicTrs, atomicTrsErr := rep.signedAtomicTrsRepository.Retrieve(atomicTrsPath)
	if atomicTrsErr == nil {
		aggrTrsBuilder.WithAtomicTransactions(atomicTrs)
	}

	//build the aggregated transactions:
	aggrTrs, aggrTrsErr := aggrTrsBuilder.Now()
	if aggrTrsErr != nil {
		return nil, aggrTrsErr
	}

	return aggrTrs, nil
}
