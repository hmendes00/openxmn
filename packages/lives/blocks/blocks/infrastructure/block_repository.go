package infrastructure

import (
	"encoding/hex"
	"errors"
	"fmt"
	"path/filepath"

	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/domain"
	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	transactions "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/domain"
)

// BlockRepository represents a concrete Block repository
type BlockRepository struct {
	blkBuilderFactory blocks.BlockBuilderFactory
	htRepository      hashtrees.HashTreeRepository
	aggTrsRepository  transactions.SignedTransactionsRepository
	objRepository     objects.ObjectRepository
}

// CreateBlockRepository creates a BlockRepository instance
func CreateBlockRepository(
	blkBuilderFactory blocks.BlockBuilderFactory,
	htRepository hashtrees.HashTreeRepository,
	aggTrsRepository transactions.SignedTransactionsRepository,
	objRepository objects.ObjectRepository,
) blocks.BlockRepository {
	out := BlockRepository{
		blkBuilderFactory: blkBuilderFactory,
		htRepository:      htRepository,
		aggTrsRepository:  aggTrsRepository,
		objRepository:     objRepository,
	}
	return &out
}

// Retrieve retrieves a Block
func (rep *BlockRepository) Retrieve(dirPath string) (blocks.Block, error) {
	//retrieve the object:
	obj, objErr := rep.objRepository.Retrieve(dirPath)
	if objErr != nil {
		return nil, objErr
	}

	//retrieve the hashtree:
	ht, htErr := rep.htRepository.Retrieve(dirPath)
	if htErr != nil {
		return nil, htErr
	}

	//retrieve the transactions:
	trsDirPath := filepath.Join(dirPath, "aggregated_transactions")
	trs, trsErr := rep.aggTrsRepository.RetrieveAll(trsDirPath)
	if trsErr != nil {
		return nil, trsErr
	}

	//create the trs map:
	trsMap := map[string]transactions.SignedTransactions{}
	for _, oneTrs := range trs {
		IDAsString := hex.EncodeToString(oneTrs.GetID().Bytes())
		trsMap[IDAsString] = oneTrs
	}

	//create the blocks:
	htBlocks := [][]byte{}
	for _, oneTrs := range trs {
		htBlocks = append(htBlocks, oneTrs.GetID().Bytes())
	}

	//re-order the blocks:
	reOrderedBlks, reOrderedBlksErr := ht.Order(htBlocks)
	if reOrderedBlksErr != nil {
		return nil, reOrderedBlksErr
	}

	//re-order the transactions:
	reorderedTrs := []transactions.SignedTransactions{}
	for _, oneBlk := range reOrderedBlks {
		IDAsString := hex.EncodeToString(oneBlk)
		if foundTrs, ok := trsMap[IDAsString]; ok {
			reorderedTrs = append(reorderedTrs, foundTrs)
			continue
		}

		str := fmt.Sprintf("there is 1 aggregated transaction (ID: %s) that was not in the HashTree", IDAsString)
		return nil, errors.New(str)
	}

	//build the block:
	met := obj.GetMetaData()
	id := met.GetID()
	ts := met.CreatedOn()
	blk, blkErr := rep.blkBuilderFactory.Create().Create().WithID(id).CreatedOn(ts).WithTransactions(reorderedTrs).Now()
	if blkErr != nil {
		return nil, blkErr
	}

	return blk, nil
}
