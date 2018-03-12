package signed

import (
	"encoding/hex"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"

	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
	signed_trs "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/transactions/signed"
)

// TransactionsRepository represents a concrete TransactionsRepository implementation
type TransactionsRepository struct {
	metaDataRepository metadata.MetaDataRepository
	trsRepository      signed_trs.TransactionRepository
	trsBuilderFactory  signed_trs.TransactionsBuilderFactory
}

// CreateTransactionsRepository creates a new TransactionsRepository instance
func CreateTransactionsRepository(metaDataRepository metadata.MetaDataRepository, trsRepository signed_trs.TransactionRepository, trsBuilderFactory signed_trs.TransactionsBuilderFactory) signed_trs.TransactionsRepository {
	out := TransactionsRepository{
		metaDataRepository: metaDataRepository,
		trsRepository:      trsRepository,
		trsBuilderFactory:  trsBuilderFactory,
	}

	return &out
}

// Retrieve retrieve a Transactions instance
func (rep *TransactionsRepository) Retrieve(dirPath string) (signed_trs.Transactions, error) {
	//retrieve the metadata:
	met, metErr := rep.metaDataRepository.Retrieve(dirPath)
	if metErr != nil {
		return nil, metErr
	}

	//retrieve the transactions:
	trsPath := filepath.Join(dirPath, "signed_transactions")
	trs, trsErr := rep.trsRepository.RetrieveAll(trsPath)
	if trsErr != nil {
		return nil, trsErr
	}

	//retrieve the blocks:
	blocks := [][]byte{
		met.GetID().Bytes(),
		[]byte(strconv.Itoa(int(met.CreatedOn().UnixNano()))),
	}

	trsMap := map[string]signed_trs.Transaction{}
	for _, oneTrs := range trs {
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
	reOrderedTrs := []signed_trs.Transaction{}
	for _, oneBlk := range reOrderedBlks[2:] {
		blkAsString := hex.EncodeToString(oneBlk)
		if oneTrs, ok := trsMap[blkAsString]; ok {
			reOrderedTrs = append(reOrderedTrs, oneTrs)
			continue
		}

		str := fmt.Sprintf("the transaction with the hash: %s could not be found", blkAsString)
		return nil, errors.New(str)
	}

	//build the transactions:
	out, outErr := rep.trsBuilderFactory.Create().Create().WithMetaData(met).WithTransactions(reOrderedTrs).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
