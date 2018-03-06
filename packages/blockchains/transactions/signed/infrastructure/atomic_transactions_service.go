package infrastructure

import (
	"path/filepath"

	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	signed_trs "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/domain"
	stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
)

// AtomicTransactionsService represents a concrete AtomicTransactionsService implementation
type AtomicTransactionsService struct {
	metaDataService               metadata.MetaDataService
	atomicTrsService              signed_trs.AtomicTransactionService
	storedAtomicTrsBuilderFactory stored_signed_transactions.AtomicTransactionsBuilderFactory
}

// CreateAtomicTransactionsService creates a new AtomicTransactionsService instance
func CreateAtomicTransactionsService(metaDataService metadata.MetaDataService, atomicTrsService signed_trs.AtomicTransactionService, storedAtomicTrsBuilderFactory stored_signed_transactions.AtomicTransactionsBuilderFactory) signed_trs.AtomicTransactionsService {
	out := AtomicTransactionsService{
		metaDataService:               metaDataService,
		atomicTrsService:              atomicTrsService,
		storedAtomicTrsBuilderFactory: storedAtomicTrsBuilderFactory,
	}

	return &out
}

// Save saves an AtomicTransactions
func (serv *AtomicTransactionsService) Save(dirPath string, trs signed_trs.AtomicTransactions) (stored_signed_transactions.AtomicTransactions, error) {
	//save the metadata:
	met := trs.GetMetaData()
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//save the atomic transactions:
	atomicTrsList := trs.GetTransactions()
	atomicTrsPath := filepath.Join(dirPath, "atomic_transactions")
	storedAtomicTrs, storedAtomicTrsErr := serv.atomicTrsService.SaveAll(atomicTrsPath, atomicTrsList)
	if storedAtomicTrsErr != nil {
		return nil, storedAtomicTrsErr
	}

	//build the stored AtomicTransactions:
	out, outErr := serv.storedAtomicTrsBuilderFactory.Create().Create().WithMetaData(storedMet).WithTransactions(storedAtomicTrs).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
