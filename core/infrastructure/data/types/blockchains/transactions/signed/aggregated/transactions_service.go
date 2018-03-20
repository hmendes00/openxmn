package aggregated

import (
	"path/filepath"

	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/metadata"
	stored_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/transactions/signed/aggregated"
	signed "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions/signed"
	aggregated "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions/signed/aggregated"
)

// TransactionsService represents a concrete TransactionsService implementation
type TransactionsService struct {
	metaDataService         metadata.MetaDataService
	signedTrsService        signed.TransactionsService
	atomicSignedTrsService  signed.AtomicTransactionsService
	storedTrsBuilderFactory stored_aggregated_transactions.TransactionsBuilderFactory
}

// CreateTransactionsService creates a new TransactionsService instance
func CreateTransactionsService(
	metaDataService metadata.MetaDataService,
	signedTrsService signed.TransactionsService,
	atomicSignedTrsService signed.AtomicTransactionsService,
	storedTrsBuilderFactory stored_aggregated_transactions.TransactionsBuilderFactory,
) aggregated.TransactionsService {
	out := TransactionsService{
		metaDataService:         metaDataService,
		signedTrsService:        signedTrsService,
		atomicSignedTrsService:  atomicSignedTrsService,
		storedTrsBuilderFactory: storedTrsBuilderFactory,
	}

	return &out
}

// Save saves an aggregated Transactions instance
func (serv *TransactionsService) Save(dirPath string, trs aggregated.Transactions) (stored_aggregated_transactions.Transactions, error) {
	//save the metadata:
	met := trs.GetMetaData()
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//create the stored aggregated transactions builder:
	aggregatedTrsBuilder := serv.storedTrsBuilderFactory.Create().Create().WithMetaData(storedMet)

	//save the transactions, if any:
	if trs.HasTransactions() {
		subTrs := trs.GetTransactions()
		trsPath := filepath.Join(dirPath, "signed_transactions")
		storedTrs, storedTrsErr := serv.signedTrsService.Save(trsPath, subTrs)
		if storedTrsErr != nil {
			return nil, storedTrsErr
		}

		//add the stored transaction to the builder:
		aggregatedTrsBuilder.WithTransactions(storedTrs)
	}

	//save the atomic transactions, if any:
	if trs.HasAtomicTransactions() {
		atomicTrs := trs.GetAtomicTransactions()
		atomicTrsPath := filepath.Join(dirPath, "signed_atomic_transactions")
		storedAtomicTrs, storedAtomicTrsErr := serv.atomicSignedTrsService.Save(atomicTrsPath, atomicTrs)
		if storedAtomicTrsErr != nil {
			return nil, storedAtomicTrsErr
		}

		//add the stored atomic transaction to the builder:
		aggregatedTrsBuilder.WithAtomicTransactions(storedAtomicTrs)
	}

	//build the stored aggregated transactions:
	aggregatedTrs, aggregatedTrsErr := aggregatedTrsBuilder.Now()
	if aggregatedTrsErr != nil {
		return nil, aggregatedTrsErr
	}

	return aggregatedTrs, nil
}
