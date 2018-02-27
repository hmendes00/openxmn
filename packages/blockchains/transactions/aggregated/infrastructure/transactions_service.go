package infrastructure

import (
	"path/filepath"

	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
	signed "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/domain"
	stored_aggregated_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/aggregated/domain"
)

// TransactionsService represents a concrete TransactionsService implementation
type TransactionsService struct {
	metaDataBuilderFactory  metadata.MetaDataBuilderFactory
	metaDataService         metadata.MetaDataService
	htService               hashtrees.HashTreeService
	signedTrsService        signed.TransactionService
	atomicSignedTrsService  signed.AtomicTransactionService
	storedTrsBuilderFactory stored_aggregated_transactions.TransactionsBuilderFactory
}

// CreateTransactionsService creates a new TransactionsService instance
func CreateTransactionsService(
	metaDataBuilderFactory metadata.MetaDataBuilderFactory,
	metaDataService metadata.MetaDataService,
	htService hashtrees.HashTreeService,
	signedTrsService signed.TransactionService,
	atomicSignedTrsService signed.AtomicTransactionService,
	storedTrsBuilderFactory stored_aggregated_transactions.TransactionsBuilderFactory,
) aggregated.TransactionsService {
	out := TransactionsService{
		metaDataBuilderFactory:  metaDataBuilderFactory,
		metaDataService:         metaDataService,
		htService:               htService,
		signedTrsService:        signedTrsService,
		atomicSignedTrsService:  atomicSignedTrsService,
		storedTrsBuilderFactory: storedTrsBuilderFactory,
	}

	return &out
}

// Save saves an aggregated Transactions instance
func (serv *TransactionsService) Save(dirPath string, trs aggregated.Transactions) (stored_aggregated_transactions.Transactions, error) {
	//build the metadata:
	id := trs.GetID()
	createdOn := trs.CreatedOn()
	metaData, metaDataErr := serv.metaDataBuilderFactory.Create().Create().WithID(id).CreatedOn(createdOn).Now()
	if metaDataErr != nil {
		return nil, metaDataErr
	}

	//save the metadata:
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, metaData)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//save the hashtree:
	ht := trs.GetHashTree()
	storedHt, storedHtErr := serv.htService.Save(dirPath, ht)
	if storedHtErr != nil {
		return nil, storedHtErr
	}

	//create the stored aggregated transactions builder:
	aggregatedTrsBuilder := serv.storedTrsBuilderFactory.Create().Create().WithMetaData(storedMet).WithHashTree(storedHt)

	//save the transactions, if any:
	if trs.HasTrs() {
		subTrs := trs.GetTrs()
		signedTrsPath := filepath.Join(dirPath, "signed_transactions")
		storedTrs, storedTrsErr := serv.signedTrsService.SaveAll(signedTrsPath, subTrs)
		if storedTrsErr != nil {
			return nil, storedTrsErr
		}

		//add the stored transaction to the builder:
		aggregatedTrsBuilder.WithTrs(storedTrs)
	}

	//save the atomic transactions, if any:
	if trs.HasAtomicTrs() {
		atomicTrs := trs.GetAtomicTrs()
		atomicTrsPath := filepath.Join(dirPath, "signed_atomic_transactions")
		storedAtomicTrs, storedAtomicTrsErr := serv.atomicSignedTrsService.SaveAll(atomicTrsPath, atomicTrs)
		if storedAtomicTrsErr != nil {
			return nil, storedAtomicTrsErr
		}

		//add the stored atomic transaction to the builder:
		aggregatedTrsBuilder.WithAtomicTrs(storedAtomicTrs)
	}

	//build the stored aggregated transactions:
	aggregatedTrs, aggregatedTrsErr := aggregatedTrsBuilder.Now()
	if aggregatedTrsErr != nil {
		return nil, aggregatedTrsErr
	}

	return aggregatedTrs, nil
}
