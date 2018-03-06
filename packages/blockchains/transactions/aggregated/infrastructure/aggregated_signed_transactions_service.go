package infrastructure

import (
	"path/filepath"

	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
	stored_aggregated_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/aggregated/domain"
)

// AggregatedSignedTransactionsService represents a concrete AggregatedSignedTransactionsService implementation
type AggregatedSignedTransactionsService struct {
	metaDataService                   metadata.MetaDataService
	signedTrsService                  aggregated.SignedTransactionsService
	storedAggrSignedTrsBuilderFactory stored_aggregated_transactions.AggregatedSignedTransactionsBuilderFactory
}

// CreateAggregatedSignedTransactionsService creates a new AggregatedSignedTransactionsService instance
func CreateAggregatedSignedTransactionsService(metaDataService metadata.MetaDataService, signedTrsService aggregated.SignedTransactionsService, storedAggrSignedTrsBuilderFactory stored_aggregated_transactions.AggregatedSignedTransactionsBuilderFactory) aggregated.AggregatedSignedTransactionsService {
	out := AggregatedSignedTransactionsService{
		metaDataService:                   metaDataService,
		signedTrsService:                  signedTrsService,
		storedAggrSignedTrsBuilderFactory: storedAggrSignedTrsBuilderFactory,
	}

	return &out
}

// Save saves an AggregatedSignedTransactions instance
func (serv *AggregatedSignedTransactionsService) Save(dirPath string, trs aggregated.AggregatedSignedTransactions) (stored_aggregated_transactions.AggregatedSignedTransactions, error) {
	//save the metaData
	met := trs.GetMetaData()
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//save the transactions:
	signedTrs := trs.GetTransactions()
	trsPath := filepath.Join(dirPath, "signed_transactions")
	storedTrs, storedTrsErr := serv.signedTrsService.SaveAll(trsPath, signedTrs)
	if storedTrsErr != nil {
		return nil, storedTrsErr
	}

	//build the aggregated signed transactions:
	out, outErr := serv.storedAggrSignedTrsBuilderFactory.Create().Create().WithMetaData(storedMet).WithTransactions(storedTrs).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
