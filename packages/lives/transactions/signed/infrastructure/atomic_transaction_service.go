package infrastructure

import (
	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	signed_trs "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	trs "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain"
	stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// AtomicTransactionService represents a concrete AtomicTransactionService implementation
type AtomicTransactionService struct {
	treeBuilderFactory stored_objects.TreeBuilderFactory
	trsService         trs.TransactionService
	objService         objects.ObjectService
	objBuilderFactory  objects.ObjectBuilderFactory
}

// CreateAtomicTransactionService creates a new AtomicTransactionService instance
func CreateAtomicTransactionService(treeBuilderFactory stored_objects.TreeBuilderFactory, trsService trs.TransactionService, objService objects.ObjectService, objBuilderFactory objects.ObjectBuilderFactory) signed_trs.AtomicTransactionService {
	out := AtomicTransactionService{
		treeBuilderFactory: treeBuilderFactory,
		trsService:         trsService,
		objService:         objService,
		objBuilderFactory:  objBuilderFactory,
	}
	return &out
}

// Save save a signed AtomicTransaction on disk
func (serv *AtomicTransactionService) Save(dirPath string, atomicTrs signed_trs.AtomicTransaction) (stored_objects.Tree, error) {

	//save the transactions:
	trs := atomicTrs.GetTrs()
	storedAtomicTrsObj, storedAtomicTrsObjErr := serv.trsService.SaveAll(dirPath, trs)
	if storedAtomicTrsObjErr != nil {
		return nil, storedAtomicTrsObjErr
	}

	//build the object:
	id := atomicTrs.GetID()
	sig := atomicTrs.GetSignature()
	createdOn := atomicTrs.CreatedOn()
	obj, objErr := serv.objBuilderFactory.Create().Create().WithID(id).WithSignature(sig).CreatedOn(createdOn).Now()
	if objErr != nil {
		return nil, objErr
	}

	//store the object:
	storedObj, storedObjErr := serv.objService.Save(dirPath, obj)
	if storedObjErr != nil {
		return nil, storedObjErr
	}

	//create the tree:
	tr, trErr := serv.treeBuilderFactory.Create().Create().WithName("atomic_transaction").WithObject(storedObj).WithSubObjects(storedAtomicTrsObj).Now()
	if trErr != nil {
		return nil, trErr
	}

	return tr, nil
}

// SaveAll saves []AtomicTransaction on disk
func (serv *AtomicTransactionService) SaveAll(dirPath string, atomicTrs []signed_trs.AtomicTransaction) ([]stored_objects.Tree, error) {
	out := []stored_objects.Tree{}
	for _, oneAtomicTrs := range atomicTrs {
		oneTree, oneTreeErr := serv.Save(dirPath, oneAtomicTrs)
		if oneTreeErr != nil {
			return nil, oneTreeErr
		}

		out = append(out, oneTree)
	}

	return out, nil
}
