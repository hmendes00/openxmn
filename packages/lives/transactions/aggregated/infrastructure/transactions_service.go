package infrastructure

import (
	"path/filepath"

	hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/domain"
	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/domain"
	signed "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// TransactionsService represents a concrete TransactionsService implementation
type TransactionsService struct {
	signedTrsService       signed.TransactionService
	atomicSignedTrsService signed.AtomicTransactionService
	htService              hashtrees.HashTreeService
	metaDataBuilderFactory objects.MetaDataBuilderFactory
	objService             objects.ObjectService
	objBuilderFactory      objects.ObjectBuilderFactory
	treeBuilderFactory     stored_objects.TreeBuilderFactory
	treesBuilderFactory    stored_objects.TreesBuilderFactory
}

// CreateTransactionsService creates a new TransactionsService instance
func CreateTransactionsService(
	signedTrsService signed.TransactionService,
	atomicSignedTrsService signed.AtomicTransactionService,
	htService hashtrees.HashTreeService,
	metaDataBuilderFactory objects.MetaDataBuilderFactory,
	objService objects.ObjectService,
	objBuilderFactory objects.ObjectBuilderFactory,
	treeBuilderFactory stored_objects.TreeBuilderFactory,
	treesBuilderFactory stored_objects.TreesBuilderFactory,
) aggregated.TransactionsService {
	out := TransactionsService{
		signedTrsService:       signedTrsService,
		atomicSignedTrsService: atomicSignedTrsService,
		htService:              htService,
		metaDataBuilderFactory: metaDataBuilderFactory,
		objService:             objService,
		objBuilderFactory:      objBuilderFactory,
		treeBuilderFactory:     treeBuilderFactory,
		treesBuilderFactory:    treesBuilderFactory,
	}

	return &out
}

// Save saves an aggregated Transactions instance
func (serv *TransactionsService) Save(dirPath string, trs aggregated.Transactions) (stored_objects.Tree, error) {
	//save the hashtree:
	ht := trs.GetHashTree()
	storedHt, storedHtErr := serv.htService.Save(dirPath, ht)
	if storedHtErr != nil {
		return nil, storedHtErr
	}

	//build the metadata:
	id := trs.GetID()
	createdOn := trs.CreatedOn()
	metaData, metaDataErr := serv.metaDataBuilderFactory.Create().Create().WithID(id).CreatedOn(createdOn).Now()
	if metaDataErr != nil {
		return nil, metaDataErr
	}

	//build the object:
	obj, objErr := serv.objBuilderFactory.Create().Create().WithMetaData(metaData).Now()
	if objErr != nil {
		return nil, objErr
	}

	//save the object:
	storedObj, storedObjErr := serv.objService.Save(dirPath, obj)
	if storedObjErr != nil {
		return nil, storedObjErr
	}

	//create the sub trees:
	storedSubTrees := []stored_objects.Tree{}

	//save the transactions, if any:
	if trs.HasTrs() {
		subTrs := trs.GetTrs()
		signedTrsPath := filepath.Join(dirPath, "signed_transactions")
		storedTrs, storedTrsErr := serv.signedTrsService.SaveAll(signedTrsPath, subTrs)
		if storedTrsErr != nil {
			return nil, storedTrsErr
		}

		//add the stored trs:
		for _, oneStoredTree := range storedTrs {
			storedSubTrees = append(storedSubTrees, oneStoredTree)
		}
	}

	//save the atomic transactions, if any:
	if trs.HasAtomicTrs() {
		atomicTrs := trs.GetAtomicTrs()
		atomicTrsPath := filepath.Join(dirPath, "signed_atomic_transactions")
		storedAtomicTrs, storedAtomicTrsErr := serv.atomicSignedTrsService.SaveAll(atomicTrsPath, atomicTrs)
		if storedAtomicTrsErr != nil {
			return nil, storedAtomicTrsErr
		}

		//add the stored atomic trs:
		for _, oneStoredTree := range storedAtomicTrs {
			storedSubTrees = append(storedSubTrees, oneStoredTree)
		}
	}

	//build the trees:
	tres, tresErr := serv.treesBuilderFactory.Create().Create().WithTrees(storedSubTrees).WithHastTree(storedHt).Now()
	if tresErr != nil {
		return nil, tresErr
	}

	//build the tree:
	tree, treeErr := serv.treeBuilderFactory.Create().Create().WithName("aggregated_transactions").WithObject(storedObj).WithSubTrees(tres).Now()
	if treeErr != nil {
		return nil, treeErr
	}

	return tree, nil
}
