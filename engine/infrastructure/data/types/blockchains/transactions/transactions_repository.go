package transactions

import (
	"encoding/hex"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"

	metadata "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/metadata"
	transactions "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/transactions"
)

// TransactionsRepository represents a concrete TransactionsRepository implementation
type TransactionsRepository struct {
	metaDataRepository metadata.MetaDataRepository
	trsRepository      transactions.TransactionRepository
	trsBuilderFactory  transactions.TransactionsBuilderFactory
}

// CreateTransactionsRepository creates a new TransactionsRepository instance
func CreateTransactionsRepository(metaDataRepository metadata.MetaDataRepository, trsRepository transactions.TransactionRepository, trsBuilderFactory transactions.TransactionsBuilderFactory) transactions.TransactionsRepository {
	out := TransactionsRepository{
		metaDataRepository: metaDataRepository,
		trsRepository:      trsRepository,
		trsBuilderFactory:  trsBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a Transactions instance
func (rep *TransactionsRepository) Retrieve(dirPath string) (transactions.Transactions, error) {
	//retrieve the metadata:
	met, metErr := rep.metaDataRepository.Retrieve(dirPath)
	if metErr != nil {
		return nil, metErr
	}

	//retrieve the transactions:
	trsPath := filepath.Join(dirPath, "transactions")
	trsList, trsListErr := rep.trsRepository.RetrieveAll(trsPath)
	if trsListErr != nil {
		return nil, trsListErr
	}

	//retrieve the blocks:
	blocks := [][]byte{
		met.GetID().Bytes(),
		[]byte(strconv.Itoa(int(met.CreatedOn().UnixNano()))),
	}

	trsMap := map[string]transactions.Transaction{}
	for _, oneTrs := range trsList {
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
	reOrderedTrs := []transactions.Transaction{}
	for _, oneBlk := range reOrderedBlks[2:] {
		blkAsString := hex.EncodeToString(oneBlk)
		if oneTrs, ok := trsMap[blkAsString]; ok {
			reOrderedTrs = append(reOrderedTrs, oneTrs)
			continue
		}

		str := fmt.Sprintf("the transaction with the hash: %s could not be found", blkAsString)
		return nil, errors.New(str)
	}

	//build the Transactions instance:
	trs, trsErr := rep.trsBuilderFactory.Create().Create().WithMetaData(met).WithTransactions(reOrderedTrs).Now()
	if trsErr != nil {
		return nil, trsErr
	}

	return trs, nil
}
