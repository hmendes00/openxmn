package infrastructure

import (
	"path/filepath"

	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	stored_aggregated_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/aggregated/domain"
)

// SignedTransactionsService represents a concrete SignedTransactions service
type SignedTransactionsService struct {
	metadataBuilderFactory                  metadata.MetaDataBuilderFactory
	metaDataService                         metadata.MetaDataService
	userSigService                          users.SignatureService
	aggregatedTrsService                    aggregated.TransactionsService
	storedAggregatedSignedTrsBuilderFactory stored_aggregated_transactions.SignedTransactionsBuilderFactory
}

// CreateSignedTransactionsService creates a new SignedTransactionsService instance
func CreateSignedTransactionsService(
	metadataBuilderFactory metadata.MetaDataBuilderFactory,
	metaDataService metadata.MetaDataService,
	userSigService users.SignatureService,
	aggregatedTrsService aggregated.TransactionsService,
	storedAggregatedSignedTrsBuilderFactory stored_aggregated_transactions.SignedTransactionsBuilderFactory,
) aggregated.SignedTransactionsService {
	out := SignedTransactionsService{
		metadataBuilderFactory:                  metadataBuilderFactory,
		metaDataService:                         metaDataService,
		userSigService:                          userSigService,
		aggregatedTrsService:                    aggregatedTrsService,
		storedAggregatedSignedTrsBuilderFactory: storedAggregatedSignedTrsBuilderFactory,
	}
	return &out
}

// Save saves a SignedTransactions instance
func (serv *SignedTransactionsService) Save(dirPath string, signedTrs aggregated.SignedTransactions) (stored_aggregated_transactions.SignedTransactions, error) {
	//build the metadata instance:
	id := signedTrs.GetID()
	ts := signedTrs.CreatedOn()
	met, metErr := serv.metadataBuilderFactory.Create().Create().WithID(id).CreatedOn(ts).Now()
	if metErr != nil {
		return nil, metErr
	}

	//save the metadata:
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//save the signature:
	sig := signedTrs.GetSignature()
	storedSig, storedSigErr := serv.userSigService.Save(dirPath, sig)
	if storedSigErr != nil {
		return nil, storedSigErr
	}

	//save the transactions on disk:
	trsPath := filepath.Join(dirPath, "transactions")
	trs := signedTrs.GetTrs()
	storedTrs, storedTrsErr := serv.aggregatedTrsService.Save(trsPath, trs)
	if storedTrsErr != nil {
		return nil, storedTrsErr
	}

	//build the stored aggregated signed transactions:
	aggrTrs, aggrTrsErr := serv.storedAggregatedSignedTrsBuilderFactory.Create().Create().WithMetaData(storedMet).WithSignature(storedSig).WithTransactions(storedTrs).Now()
	if aggrTrsErr != nil {
		return nil, aggrTrsErr
	}

	return aggrTrs, nil
}

// SaveAll saves []SignedTransactions instances
func (serv *SignedTransactionsService) SaveAll(dirPath string, trs []aggregated.SignedTransactions) ([]stored_aggregated_transactions.SignedTransactions, error) {
	out := []stored_aggregated_transactions.SignedTransactions{}
	for _, oneTrs := range trs {
		trsDirPath := filepath.Join(dirPath, oneTrs.GetID().String())
		oneStoredTrs, oneStoredTrsErr := serv.Save(trsDirPath, oneTrs)
		if oneStoredTrsErr != nil {
			return nil, oneStoredTrsErr
		}

		out = append(out, oneStoredTrs)
	}

	return out, nil
}
