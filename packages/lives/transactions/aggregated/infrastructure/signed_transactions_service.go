package infrastructure

import (
	"path/filepath"

	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/domain"
	stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// SignedTransactionsService represents a concrete SignedTransactions service
type SignedTransactionsService struct {
	metadataBuilderFactory   objects.MetaDataBuilderFactory
	objBuilderFactory        objects.ObjectBuilderFactory
	objService               objects.ObjectService
	trsService               aggregated.TransactionsService
	storedTreeBuilderFactory stored_objects.TreeBuilderFactory
}

// CreateSignedTransactionsService creates a new SignedTransactionsService instance
func CreateSignedTransactionsService(
	metadataBuilderFactory objects.MetaDataBuilderFactory,
	objBuilderFactory objects.ObjectBuilderFactory,
	objService objects.ObjectService,
	trsService aggregated.TransactionsService,
	storedTreeBuilderFactory stored_objects.TreeBuilderFactory,
) aggregated.SignedTransactionsService {
	out := SignedTransactionsService{
		metadataBuilderFactory:   metadataBuilderFactory,
		objBuilderFactory:        objBuilderFactory,
		objService:               objService,
		trsService:               trsService,
		storedTreeBuilderFactory: storedTreeBuilderFactory,
	}
	return &out
}

// Save saves a SignedTransactions instance
func (serv *SignedTransactionsService) Save(dirPath string, signedTrs aggregated.SignedTransactions) (stored_objects.Tree, error) {
	//build the metadata instance:
	id := signedTrs.GetID()
	ts := signedTrs.CreatedOn()
	sig := signedTrs.GetSignature()
	met, metErr := serv.metadataBuilderFactory.Create().Create().WithID(id).CreatedOn(ts).WithSignature(sig).Now()
	if metErr != nil {
		return nil, metErr
	}

	//build the object instance:
	obj, objErr := serv.objBuilderFactory.Create().Create().WithMetaData(met).Now()
	if objErr != nil {
		return nil, objErr
	}

	//save the object on disk:
	storedObj, storedObjErr := serv.objService.Save(dirPath, obj)
	if storedObjErr != nil {
		return nil, storedObjErr
	}

	//save the transactions on disk:
	trsPath := filepath.Join(dirPath, "transactions")
	trs := signedTrs.GetTrs()
	storedTrs, storedTrsErr := serv.trsService.Save(trsPath, trs)
	if storedTrsErr != nil {
		return nil, storedTrsErr
	}

	//build the tree instance:
	tre, treErr := serv.storedTreeBuilderFactory.Create().Create().WithName("aggregated_signed_transactions").WithObject(storedObj).WithSubTree(storedTrs).Now()
	if treErr != nil {
		return nil, treErr
	}

	return tre, nil
}

// SaveAll saves []SignedTransactions instances
func (serv *SignedTransactionsService) SaveAll(dirPath string, trs []aggregated.SignedTransactions) ([]stored_objects.Tree, error) {
	return nil, nil
}
