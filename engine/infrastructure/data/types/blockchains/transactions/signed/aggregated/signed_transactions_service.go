package aggregated

import (
	"path/filepath"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	stored_aggregated_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions/signed/aggregated"
	aggregated "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions/signed/aggregated"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// SignedTransactionsService represents a concrete SignedTransactions service
type SignedTransactionsService struct {
	metaDataService                         metadata.Service
	userSigService                          users.SignatureService
	aggregatedTrsService                    aggregated.TransactionsService
	storedAggregatedSignedTrsBuilderFactory stored_aggregated_transactions.SignedTransactionsBuilderFactory
}

// CreateSignedTransactionsService creates a new SignedTransactionsService instance
func CreateSignedTransactionsService(
	metaDataService metadata.Service,
	userSigService users.SignatureService,
	aggregatedTrsService aggregated.TransactionsService,
	storedAggregatedSignedTrsBuilderFactory stored_aggregated_transactions.SignedTransactionsBuilderFactory,
) aggregated.SignedTransactionsService {
	out := SignedTransactionsService{
		metaDataService:                         metaDataService,
		userSigService:                          userSigService,
		aggregatedTrsService:                    aggregatedTrsService,
		storedAggregatedSignedTrsBuilderFactory: storedAggregatedSignedTrsBuilderFactory,
	}
	return &out
}

// Save saves a SignedTransactions instance
func (serv *SignedTransactionsService) Save(dirPath string, signedTrs aggregated.SignedTransactions) (stored_aggregated_transactions.SignedTransactions, error) {
	//save the metadata:
	met := signedTrs.GetMetaData()
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//save the signature:
	sigPath := filepath.Join(dirPath, "signature")
	sig := signedTrs.GetSignature()
	storedSig, storedSigErr := serv.userSigService.Save(sigPath, sig)
	if storedSigErr != nil {
		return nil, storedSigErr
	}

	//save the transactions on disk:
	trsPath := filepath.Join(dirPath, "transactions")
	trs := signedTrs.GetTransactions()
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
		trsDirPath := filepath.Join(dirPath, oneTrs.GetMetaData().GetID().String())
		oneStoredTrs, oneStoredTrsErr := serv.Save(trsDirPath, oneTrs)
		if oneStoredTrsErr != nil {
			return nil, oneStoredTrsErr
		}

		out = append(out, oneStoredTrs)
	}

	return out, nil
}
