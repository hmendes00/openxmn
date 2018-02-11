package infrastructure

import (
	chunks "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	trs "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain"
	stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// TransactionService represents a concrete TransactionService implementation
type TransactionService struct {
	objService              objects.ObjectService
	chkBuilderFactory       chunks.ChunksBuilderFactory
	objBuilderFactory       objects.ObjectBuilderFactory
	storedObjBuilderFactory stored_objects.ObjectBuilderFactory
}

// CreateTransactionService creates a new TransactionService instance
func CreateTransactionService(objService objects.ObjectService, chkBuilderFactory chunks.ChunksBuilderFactory, objBuilderFactory objects.ObjectBuilderFactory, storedObjBuilderFactory stored_objects.ObjectBuilderFactory) trs.TransactionService {
	out := TransactionService{
		objService:              objService,
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

	//build the object:
	id := trs.GetID()
	createdOn := trs.CreatedOn()
	obj, objErr := serv.objBuilderFactory.Create().Create().WithID(id).WithChunks(chks).CreatedOn(createdOn).Now()
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
		oneObj, oneObjErr := serv.Save(dirPath, oneTrs)
		if oneObjErr != nil {
			return nil, oneObjErr
		}

		out = append(out, oneObj)
	}

	return out, nil
}
