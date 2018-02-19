package infrastructure

import (
	"path/filepath"

	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/domain"
	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	transactions "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/domain"
	stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// BlockService represents a concrete Block service
type BlockService struct {
	storedTreesBuilderFactory stored_objects.TreesBuilderFactory
	storedTreeBuilderFactory  stored_objects.TreeBuilderFactory
	metBuilderFactory         objects.MetaDataBuilderFactory
	objBuilderFactory         objects.ObjectBuilderFactory
	objService                objects.ObjectService
	htService                 hashtrees.HashTreeService
	trsService                transactions.SignedTransactionsService
}

// CreateBlockService creates a new BlockService instance
func CreateBlockService(
	storedTreesBuilderFactory stored_objects.TreesBuilderFactory,
	storedTreeBuilderFactory stored_objects.TreeBuilderFactory,
	metBuilderFactory objects.MetaDataBuilderFactory,
	objBuilderFactory objects.ObjectBuilderFactory,
	objService objects.ObjectService,
	htService hashtrees.HashTreeService,
	trsService transactions.SignedTransactionsService,
) blocks.BlockService {
	out := BlockService{
		storedTreesBuilderFactory: storedTreesBuilderFactory,
		storedTreeBuilderFactory:  storedTreeBuilderFactory,
		metBuilderFactory:         metBuilderFactory,
		objBuilderFactory:         objBuilderFactory,
		objService:                objService,
		htService:                 htService,
		trsService:                trsService,
	}
	return &out
}

// Save saves a Block instance
func (serv *BlockService) Save(dirPath string, blk blocks.Block) (stored_objects.Tree, error) {
	//build the metadata:
	id := blk.GetID()
	ts := blk.CreatedOn()
	met, metErr := serv.metBuilderFactory.Create().Create().WithID(id).CreatedOn(ts).Now()
	if metErr != nil {
		return nil, metErr
	}

	//build the object:
	obj, objErr := serv.objBuilderFactory.Create().Create().WithMetaData(met).Now()
	if objErr != nil {
		return nil, objErr
	}

	//save the object:
	storedObj, storedObjErr := serv.objService.Save(dirPath, obj)
	if storedObjErr != nil {
		return nil, storedObjErr
	}

	//save the hashtree:
	ht := blk.GetHashTree()
	storedHt, storedHtErr := serv.htService.Save(dirPath, ht)
	if storedHtErr != nil {
		return nil, storedHtErr
	}

	//save the aggregated transactions:
	aggTrs := blk.GetTransactions()
	trsDirPath := filepath.Join(dirPath, "aggregated_transactions")
	storedTrs, storedTrsErr := serv.trsService.SaveAll(trsDirPath, aggTrs)
	if storedTrsErr != nil {
		return nil, storedTrsErr
	}

	//build the transaction trees:
	storedTrsTrees, storedTrsTreesErr := serv.storedTreesBuilderFactory.Create().Create().WithHastTree(storedHt).WithTrees(storedTrs).Now()
	if storedTrsTreesErr != nil {
		return nil, storedTrsTreesErr
	}

	//build the tree:
	storedTree, storedTreeErr := serv.storedTreeBuilderFactory.Create().Create().WithName("block").WithObject(storedObj).WithSubTrees(storedTrsTrees).Now()
	if storedTreeErr != nil {
		return nil, storedTreeErr
	}

	return storedTree, nil
}
