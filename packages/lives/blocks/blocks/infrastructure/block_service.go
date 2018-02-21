package infrastructure

import (
	"path/filepath"

	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/domain"
	metadata "github.com/XMNBlockchain/core/packages/lives/metadata/domain"
	aggregated_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/domain"
	stored_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/blocks/domain"
)

// BlockService represents a concrete Block service
type BlockService struct {
	metaDataBuilderFactory    metadata.MetaDataBuilderFactory
	metaDataService           metadata.MetaDataService
	htService                 hashtrees.HashTreeService
	aggregatedTrsService      aggregated_transactions.SignedTransactionsService
	storedBlockBuilderFactory stored_blocks.BlockBuilderFactory
}

// CreateBlockService creates a new BlockService instance
func CreateBlockService(
	metaDataBuilderFactory metadata.MetaDataBuilderFactory,
	metaDataService metadata.MetaDataService,
	htService hashtrees.HashTreeService,
	aggregatedTrsService aggregated_transactions.SignedTransactionsService,
	storedBlockBuilderFactory stored_blocks.BlockBuilderFactory,
) blocks.BlockService {
	out := BlockService{
		metaDataBuilderFactory:    metaDataBuilderFactory,
		metaDataService:           metaDataService,
		htService:                 htService,
		aggregatedTrsService:      aggregatedTrsService,
		storedBlockBuilderFactory: storedBlockBuilderFactory,
	}
	return &out
}

// Save saves a Block instance
func (serv *BlockService) Save(dirPath string, blk blocks.Block) (stored_blocks.Block, error) {
	//build the metadata:
	id := blk.GetID()
	ts := blk.CreatedOn()
	met, metErr := serv.metaDataBuilderFactory.Create().Create().WithID(id).CreatedOn(ts).Now()
	if metErr != nil {
		return nil, metErr
	}

	//save the metadata:
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
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
	storedTrs, storedTrsErr := serv.aggregatedTrsService.SaveAll(trsDirPath, aggTrs)
	if storedTrsErr != nil {
		return nil, storedTrsErr
	}

	//build the stored block:
	storedBlk, storedBlkErr := serv.storedBlockBuilderFactory.Create().Create().WithHashTree(storedHt).WithMetaData(storedMet).WithTransactions(storedTrs).Now()
	if storedBlkErr != nil {
		return nil, storedBlkErr
	}

	return storedBlk, nil
}
