package infrastructure

import (
	"encoding/json"
	"path/filepath"

	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	signed_trs "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	trs "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain"
	stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// AtomicTransactionService represents a concrete AtomicTransactionService implementation
type AtomicTransactionService struct {
	metaDataBuilderFactory   objects.MetaDataBuilderFactory
	fileBuilderFactory       files.FileBuilderFactory
	fileService              files.FileService
	treeBuilderFactory       stored_objects.TreeBuilderFactory
	trsService               trs.TransactionService
	objService               objects.ObjectService
	objBuilderFactory        objects.ObjectBuilderFactory
	storedObjsBuilderFactory stored_objects.ObjectsBuilderFactory
}

// CreateAtomicTransactionService creates a new AtomicTransactionService instance
func CreateAtomicTransactionService(
	metaDataBuilderFactory objects.MetaDataBuilderFactory,
	fileBuilderFactory files.FileBuilderFactory,
	fileService files.FileService,
	treeBuilderFactory stored_objects.TreeBuilderFactory,
	trsService trs.TransactionService,
	objService objects.ObjectService,
	objBuilderFactory objects.ObjectBuilderFactory,
	objsBuilderFactory objects.ObjectsBuilderFactory,
	storedObjsBuilderFactory stored_objects.ObjectsBuilderFactory,
) signed_trs.AtomicTransactionService {
	out := AtomicTransactionService{
		metaDataBuilderFactory:   metaDataBuilderFactory,
		fileBuilderFactory:       fileBuilderFactory,
		fileService:              fileService,
		treeBuilderFactory:       treeBuilderFactory,
		trsService:               trsService,
		objService:               objService,
		objBuilderFactory:        objBuilderFactory,
		storedObjsBuilderFactory: storedObjsBuilderFactory,
	}
	return &out
}

// Save save a signed AtomicTransaction on disk
func (serv *AtomicTransactionService) Save(dirPath string, atomicTrs signed_trs.AtomicTransaction) (stored_objects.Tree, error) {

	//save the transactions:
	trs := atomicTrs.GetTrs()
	trsDirPath := filepath.Join(dirPath, "transactions")
	storedAtomicTrsObjs, storedAtomicTrsObjsErr := serv.trsService.SaveAll(trsDirPath, trs)
	if storedAtomicTrsObjsErr != nil {
		return nil, storedAtomicTrsObjsErr
	}

	//convert the hashtree to json:
	ht := atomicTrs.GetHashTree()
	js, jsErr := json.Marshal(ht)
	if jsErr != nil {
		return nil, jsErr
	}

	//build the hashtree file:
	htFile, htFileErr := serv.fileBuilderFactory.Create().Create().WithData(js).WithFileName("hashtree").WithExtension("json").Now()
	if htFileErr != nil {
		return nil, htFileErr
	}

	//save the hashtree:
	storedHt, storedHtErr := serv.fileService.Save(dirPath, htFile)
	if storedHtErr != nil {
		return nil, storedHtErr
	}

	//build the stored objects:
	storedObjects, storedObjectsErr := serv.storedObjsBuilderFactory.Create().Create().WithObjects(storedAtomicTrsObjs).WithHashTree(storedHt).Now()
	if storedObjectsErr != nil {
		return nil, storedObjectsErr
	}

	//build the metadata:
	id := atomicTrs.GetID()
	sig := atomicTrs.GetSignature()
	createdOn := atomicTrs.CreatedOn()
	met, metErr := serv.metaDataBuilderFactory.Create().Create().WithID(id).WithSignature(sig).CreatedOn(createdOn).Now()
	if metErr != nil {
		return nil, metErr
	}

	//build the object:
	obj, objErr := serv.objBuilderFactory.Create().Create().WithMetaData(met).Now()
	if objErr != nil {
		return nil, objErr
	}

	//store the object:
	storedObj, storedObjErr := serv.objService.Save(dirPath, obj)
	if storedObjErr != nil {
		return nil, storedObjErr
	}

	//create the tree:
	tr, trErr := serv.treeBuilderFactory.Create().Create().WithName("atomic_transaction").WithObject(storedObj).WithSubObjects(storedObjects).Now()
	if trErr != nil {
		return nil, trErr
	}

	return tr, nil
}

// SaveAll saves []AtomicTransaction on disk
func (serv *AtomicTransactionService) SaveAll(dirPath string, atomicTrs []signed_trs.AtomicTransaction) ([]stored_objects.Tree, error) {
	out := []stored_objects.Tree{}
	for _, oneAtomicTrs := range atomicTrs {
		oneAtomicTrsDirPath := filepath.Join(dirPath, oneAtomicTrs.GetID().String())
		oneTree, oneTreeErr := serv.Save(oneAtomicTrsDirPath, oneAtomicTrs)
		if oneTreeErr != nil {
			return nil, oneTreeErr
		}

		out = append(out, oneTree)
	}

	return out, nil
}
