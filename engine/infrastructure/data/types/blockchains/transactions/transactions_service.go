package transactions

import (
	"path/filepath"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	stored_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
)

// TransactionsService represents a concrete Transactions Service implementation
type TransactionsService struct {
	metaDataService         metadata.MetaDataService
	trsService              transactions.TransactionService
	storedTrsBuilderFactory stored_transactions.BuilderFactory
}

// CreateTransactionsService creates a new TransactionsService
func CreateTransactionsService(metaDataService metadata.MetaDataService, trsService transactions.TransactionService, storedTrsBuilderFactory stored_transactions.BuilderFactory) transactions.TransactionsService {
	out := TransactionsService{
		metaDataService:         metaDataService,
		trsService:              trsService,
		storedTrsBuilderFactory: storedTrsBuilderFactory,
	}

	return &out
}

// Save saves a Transactions
func (serv *TransactionsService) Save(dirPath string, trs transactions.Transactions) (stored_transactions.Transactions, error) {
	//save the metadata:
	met := trs.GetMetaData()
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//save the transactions:
	trsList := trs.GetTransactions()
	trsPath := filepath.Join(dirPath, "transactions")
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
