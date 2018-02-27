package infrastructure

import (
	"path/filepath"

	chunks "github.com/XMNBlockchain/core/packages/blockchains/chunks/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	trs "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
	stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/domain"
)

// TransactionService represents a concrete TransactionService implementation
type TransactionService struct {
	metaDataBuilderFactory  metadata.MetaDataBuilderFactory
	metaDataService         metadata.MetaDataService
	chkBuilderFactory       chunks.ChunksBuilderFactory
	chkService              chunks.ChunksService
	storedTrsBuilderFactory stored_transactions.TransactionBuilderFactory
}

// CreateTransactionService creates a new TransactionService instance
func CreateTransactionService(
	metaDataBuilderFactory metadata.MetaDataBuilderFactory,
	metaDataService metadata.MetaDataService,
	chkBuilderFactory chunks.ChunksBuilderFactory,
	chkService chunks.ChunksService,
	storedTrsBuilderFactory stored_transactions.TransactionBuilderFactory,
) trs.TransactionService {
	out := TransactionService{
		metaDataBuilderFactory:  metaDataBuilderFactory,
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
	chks, chksErr := serv.chkBuilderFactory.Create().Create().WithInstance(trs).Now()
	if chksErr != nil {
		return nil, chksErr
	}

	//save the chunks:
	storedChks, storedChksErr := serv.chkService.Save(dirPath, chks)
	if storedChksErr != nil {
		return nil, storedChksErr
	}

	//build the metaData:
	id := trs.GetID()
	createdOn := trs.CreatedOn()
	met, metErr := serv.metaDataBuilderFactory.Create().Create().WithID(id).CreatedOn(createdOn).Now()
	if metErr != nil {
		return nil, metErr
	}

	//save the metadata:
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
		oneObjDirPath := filepath.Join(dirPath, oneTrs.GetID().String())
		oneObj, oneObjErr := serv.Save(oneObjDirPath, oneTrs)
		if oneObjErr != nil {
			return nil, oneObjErr
		}

		out = append(out, oneObj)
	}

	return out, nil
}
