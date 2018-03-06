package infrastructure

import (
	"path/filepath"

	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	signed_trs "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/domain"
	stored_signed_transaction "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
)

// TransactionsService represents a concrete TransactionsService implementation
type TransactionsService struct {
	metaDataService         metadata.MetaDataService
	trsService              signed_trs.TransactionService
	storedTrsBuilderFactory stored_signed_transaction.TransactionsBuilderFactory
}

// CreateTransactionsService creates a new TransactionsService instance
func CreateTransactionsService(metaDataService metadata.MetaDataService, trsService signed_trs.TransactionService, storedTrsBuilderFactory stored_signed_transaction.TransactionsBuilderFactory) signed_trs.TransactionsService {
	out := TransactionsService{
		metaDataService:         metaDataService,
		trsService:              trsService,
		storedTrsBuilderFactory: storedTrsBuilderFactory,
	}

	return &out
}

// Save save Transactions
func (serv *TransactionsService) Save(dirPath string, trs signed_trs.Transactions) (stored_signed_transaction.Transactions, error) {
	//save the metadata:
	met := trs.GetMetaData()
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//save the transactions:
	trsList := trs.GetTransactions()
	trsPath := filepath.Join(dirPath, "signed_transactions")
	storedTrs, storedTrsErr := serv.trsService.SaveAll(trsPath, trsList)
	if storedTrsErr != nil {
		return nil, storedTrsErr
	}

	//build the stored transactions:
	out, outErr := serv.storedTrsBuilderFactory.Create().Create().WithMetaData(storedMet).WithTransactions(storedTrs).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
