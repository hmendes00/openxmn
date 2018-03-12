package blocks

import (
	"encoding/hex"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"

	blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/blocks"
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
	aggregated "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/transactions/signed/aggregated"
)

// BlockRepository represents a concrete BlockRepository implementation
type BlockRepository struct {
	metaDataRepository  metadata.MetaDataRepository
	signedTrsRepository aggregated.SignedTransactionsRepository
	blkBuilderFactory   blocks.BlockBuilderFactory
}

// CreateBlockRepository creates a new BlockRepository instance
func CreateBlockRepository(metaDataRepository metadata.MetaDataRepository, signedTrsRepository aggregated.SignedTransactionsRepository, blkBuilderFactory blocks.BlockBuilderFactory) blocks.BlockRepository {
	out := BlockRepository{
		metaDataRepository:  metaDataRepository,
		signedTrsRepository: signedTrsRepository,
		blkBuilderFactory:   blkBuilderFactory,
	}

	return &out
}

// Retrieve retrieves an Block instance
func (rep *BlockRepository) Retrieve(dirPath string) (blocks.Block, error) {
	//retrieve the metadata:
	met, metErr := rep.metaDataRepository.Retrieve(dirPath)
	if metErr != nil {
		return nil, metErr
	}

	//retrieve the transactions:
	trsPath := filepath.Join(dirPath, "signed_transactions")
	signedTrs, signedTrsErr := rep.signedTrsRepository.RetrieveAll(trsPath)
	if signedTrsErr != nil {
		return nil, signedTrsErr
	}

	//retrieve the blocks:
	blocks := [][]byte{
		met.GetID().Bytes(),
		[]byte(strconv.Itoa(int(met.CreatedOn().UnixNano()))),
	}

	trsMap := map[string]aggregated.SignedTransactions{}
	for _, oneTrs := range signedTrs {
		hash := oneTrs.GetMetaData().GetHashTree().GetHash()
		trsMap[hash.String()] = oneTrs
		blocks = append(blocks, hash.Get())
	}

	//re-order the blocks:
	reOrderedBlks, reOrderedBlksErr := met.GetHashTree().Order(blocks)
	if reOrderedBlksErr != nil {
		return nil, reOrderedBlksErr
	}

	//re-order the transactions:
	reOrderedTrs := []aggregated.SignedTransactions{}
	for _, oneBlk := range reOrderedBlks[2:] {
		blkAsString := hex.EncodeToString(oneBlk)
		if oneTrs, ok := trsMap[blkAsString]; ok {
			reOrderedTrs = append(reOrderedTrs, oneTrs)
			continue
		}

		str := fmt.Sprintf("the signed transactions with the hash: %s could not be found", blkAsString)
		return nil, errors.New(str)
	}

	//build the aggregated signed transactions:
	out, outErr := rep.blkBuilderFactory.Create().Create().WithMetaData(met).WithTransactions(reOrderedTrs).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
