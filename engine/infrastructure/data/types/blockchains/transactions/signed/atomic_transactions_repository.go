package signed

import (
	"encoding/hex"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"

	metadata "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/metadata"
	signed_trs "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/transactions/signed"
)

// AtomicTransactionsRepository represents a concrete AtomicTransactionsRepostory implementation
type AtomicTransactionsRepository struct {
	metaDataRepository        metadata.MetaDataRepository
	atomicTrsRepository       signed_trs.AtomicTransactionRepository
	atomicTransBuilderFactory signed_trs.AtomicTransactionsBuilderFactory
}

// CreateAtomicTransactionsRepository creates a new AtomicTransactionsRepository instance
func CreateAtomicTransactionsRepository(metaDataRepository metadata.MetaDataRepository, atomicTrsRepository signed_trs.AtomicTransactionRepository, atomicTransBuilderFactory signed_trs.AtomicTransactionsBuilderFactory) signed_trs.AtomicTransactionsRepository {
	out := AtomicTransactionsRepository{
		metaDataRepository:        metaDataRepository,
		atomicTrsRepository:       atomicTrsRepository,
		atomicTransBuilderFactory: atomicTransBuilderFactory,
	}

	return &out
}

// Retrieve retrieves an AtomicTransactions instance
func (rep *AtomicTransactionsRepository) Retrieve(dirPath string) (signed_trs.AtomicTransactions, error) {
	//retrieve the metadata:
	met, metErr := rep.metaDataRepository.Retrieve(dirPath)
	if metErr != nil {
		return nil, metErr
	}

	//retrieve the atomic transactions:
	atomicTrsPath := filepath.Join(dirPath, "atomic_transactions")
	atomicTrs, atomicTrsErr := rep.atomicTrsRepository.RetrieveAll(atomicTrsPath)
	if atomicTrsErr != nil {
		return nil, atomicTrsErr
	}

	//retrieve the blocks:
	blocks := [][]byte{
		met.GetID().Bytes(),
		[]byte(strconv.Itoa(int(met.CreatedOn().UnixNano()))),
	}

	trsMap := map[string]signed_trs.AtomicTransaction{}
	for _, oneTrs := range atomicTrs {
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
	reOrderedTrs := []signed_trs.AtomicTransaction{}
	for _, oneBlk := range reOrderedBlks[2:] {
		blkAsString := hex.EncodeToString(oneBlk)
		if oneTrs, ok := trsMap[blkAsString]; ok {
			reOrderedTrs = append(reOrderedTrs, oneTrs)
			continue
		}

		str := fmt.Sprintf("the atomic transaction with the hash: %s could not be found", blkAsString)
		return nil, errors.New(str)
	}

	//build the AtomicTransactions instance:
	out, outErr := rep.atomicTransBuilderFactory.Create().Create().WithMetaData(met).WithTransactions(reOrderedTrs).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
