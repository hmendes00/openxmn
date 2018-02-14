package infrastructure

import (
	"path/filepath"

	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	signed_trs "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	trs "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain"
	stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// TransactionService represents a concrete TransactionService implementation
type TransactionService struct {
	metaDataBuilderFactory objects.MetaDataBuilderFactory
	treeBuilderFactory     stored_objects.TreeBuilderFactory
	trsService             trs.TransactionService
	objService             objects.ObjectService
	objBuilderFactory      objects.ObjectBuilderFactory
}

// CreateTransactionService creates a new TransactionService instance
func CreateTransactionService(
	metaDataBuilderFactory objects.MetaDataBuilderFactory,
	treeBuilderFactory stored_objects.TreeBuilderFactory,
	trsService trs.TransactionService,
	objService objects.ObjectService,
	objBuilderFactory objects.ObjectBuilderFactory,
) signed_trs.TransactionService {
	out := TransactionService{
		metaDataBuilderFactory: metaDataBuilderFactory,
		treeBuilderFactory:     treeBuilderFactory,
		trsService:             trsService,
		objService:             objService,
		objBuilderFactory:      objBuilderFactory,
	}
	return &out
}

// Save save a signed Transaction on disk
func (serv *TransactionService) Save(dirPath string, signedTrs signed_trs.Transaction) (stored_objects.Tree, error) {

	//save the transaction:
	trs := signedTrs.GetTrs()
	trsPath := filepath.Join(dirPath, "transaction")
	storedTrsObj, storedTrsObjErr := serv.trsService.Save(trsPath, trs)
	if storedTrsObjErr != nil {
		return nil, storedTrsObjErr
	}

	//build the metadata:
	id := signedTrs.GetID()
	sig := signedTrs.GetSignature()
	createdOn := signedTrs.CreatedOn()
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
	tr, trErr := serv.treeBuilderFactory.Create().Create().WithName("signed_transactions").WithObject(storedObj).WithSubObject(storedTrsObj).Now()
	if trErr != nil {
		return nil, trErr
	}

	return tr, nil
}

// SaveAll saves signed []Transaction on disk
func (serv *TransactionService) SaveAll(dirPath string, signedTrs []signed_trs.Transaction) ([]stored_objects.Tree, error) {
	out := []stored_objects.Tree{}
	for _, oneSignedTrs := range signedTrs {
		signedTrsPath := filepath.Join(dirPath, oneSignedTrs.GetID().String())
		oneTree, oneTreeErr := serv.Save(signedTrsPath, oneSignedTrs)
		if oneTreeErr != nil {
			return nil, oneTreeErr
		}

		out = append(out, oneTree)
	}

	return out, nil
}
