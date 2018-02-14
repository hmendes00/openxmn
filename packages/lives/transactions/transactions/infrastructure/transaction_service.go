package infrastructure

import (
	"path/filepath"

	chunks "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	trs "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain"
	stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// TransactionService represents a concrete TransactionService implementation
type TransactionService struct {
	objService              objects.ObjectService
	metaDataBuilderFactory  objects.MetaDataBuilderFactory
	chkBuilderFactory       chunks.ChunksBuilderFactory
	objBuilderFactory       objects.ObjectBuilderFactory
	storedObjBuilderFactory stored_objects.ObjectBuilderFactory
}

// CreateTransactionService creates a new TransactionService instance
func CreateTransactionService(
	objService objects.ObjectService,
	metaDataBuilderFactory objects.MetaDataBuilderFactory,
	chkBuilderFactory chunks.ChunksBuilderFactory,
	objBuilderFactory objects.ObjectBuilderFactory,
	storedObjBuilderFactory stored_objects.ObjectBuilderFactory,
) trs.TransactionService {
	out := TransactionService{
		objService:              objService,
		metaDataBuilderFactory:  metaDataBuilderFactory,
		chkBuilderFactory:       chkBuilderFactory,
		objBuilderFactory:       objBuilderFactory,
		storedObjBuilderFactory: storedObjBuilderFactory,
	}
	return &out
}

// Save save a Transaction on disk
func (serv *TransactionService) Save(dirPath string, trs trs.Transaction) (stored_objects.Object, error) {
	//build the chunks:
	chks, chksErr := serv.chkBuilderFactory.Create().Create().WithInstance(trs).Now()
	if chksErr != nil {
		return nil, chksErr
	}

	//build the metaData:
	id := trs.GetID()
	createdOn := trs.CreatedOn()
	met, metErr := serv.metaDataBuilderFactory.Create().Create().WithID(id).CreatedOn(createdOn).Now()
	if metErr != nil {
		return nil, metErr
	}

	//build the object:
	obj, objErr := serv.objBuilderFactory.Create().Create().WithMetaData(met).WithChunks(chks).Now()
	if objErr != nil {
		return nil, objErr
	}

	//save the object:
	storedObj, storedObjErr := serv.objService.Save(dirPath, obj)
	if storedObjErr != nil {
		return nil, storedObjErr
	}

	return storedObj, nil
}

// SaveAll saves []Transaction on disk
func (serv *TransactionService) SaveAll(dirPath string, trs []trs.Transaction) ([]stored_objects.Object, error) {
	out := []stored_objects.Object{}
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
