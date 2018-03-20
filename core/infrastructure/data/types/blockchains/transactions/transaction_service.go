package transactions

import (
	"path/filepath"

	chunks "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/chunks"
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/metadata"
	stored_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/transactions"
	trs "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions"
)

// TransactionService represents a concrete TransactionService implementation
type TransactionService struct {
	metaDataService         metadata.MetaDataService
	chkBuilderFactory       chunks.BuilderFactory
	chkService              chunks.Service
	storedTrsBuilderFactory stored_transactions.TransactionBuilderFactory
}

// CreateTransactionService creates a new TransactionService instance
func CreateTransactionService(
	metaDataService metadata.MetaDataService,
	chkBuilderFactory chunks.BuilderFactory,
	chkService chunks.Service,
	storedTrsBuilderFactory stored_transactions.TransactionBuilderFactory,
) trs.TransactionService {
	out := TransactionService{
		metaDataService:         metaDataService,
		chkBuilderFactory:       chkBuilderFactory,
		chkService:              chkService,
		storedTrsBuilderFactory: storedTrsBuilderFactory,
	}
	return &out
}

// Save save a Transaction on disk
func (serv *TransactionService) Save(dirPath string, trs trs.Transaction) (stored_transactions.Transaction, error) {
	//build the chunks:
	jsData := trs.GetJSON()
	chks, chksErr := serv.chkBuilderFactory.Create().Create().WithData(jsData).Now()
	if chksErr != nil {
		return nil, chksErr
	}

	//save the chunks:
	chkDirPath := filepath.Join(dirPath, "json")
	storedChks, storedChksErr := serv.chkService.Save(chkDirPath, chks)
	if storedChksErr != nil {
		return nil, storedChksErr
	}

	//save the metadata:
	met := trs.GetMetaData()
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//build the stored transaction:
	storedTrs, storedTrsErr := serv.storedTrsBuilderFactory.Create().Create().WithMetaData(storedMet).WithChunks(storedChks).Now()
	if storedTrsErr != nil {
		return nil, storedTrsErr
	}

	return storedTrs, nil
}

// SaveAll saves []Transaction on disk
func (serv *TransactionService) SaveAll(dirPath string, trs []trs.Transaction) ([]stored_transactions.Transaction, error) {
	out := []stored_transactions.Transaction{}
	for _, oneTrs := range trs {
		oneObjDirPath := filepath.Join(dirPath, oneTrs.GetMetaData().GetID().String())
		oneObj, oneObjErr := serv.Save(oneObjDirPath, oneTrs)
		if oneObjErr != nil {
			return nil, oneObjErr
		}

		out = append(out, oneObj)
	}

	return out, nil
}
