package infrastructure

import (
	"encoding/hex"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"

	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
)

// AggregatedSignedTransactionsRepository represents a concrete AggregatedSignedTransactionsRepository implementation
type AggregatedSignedTransactionsRepository struct {
	metaDataRepository          metadata.MetaDataRepository
	signedTrsRepository         aggregated.SignedTransactionsRepository
	aggrSignedTrsBuilderFactory aggregated.AggregatedSignedTransactionsBuilderFactory
}

// CreateAggregatedSignedTransactionsRepository creates a new AggregatedSignedTransactionsRepository instance
func CreateAggregatedSignedTransactionsRepository(metaDataRepository metadata.MetaDataRepository, signedTrsRepository aggregated.SignedTransactionsRepository, aggrSignedTrsBuilderFactory aggregated.AggregatedSignedTransactionsBuilderFactory) aggregated.AggregatedSignedTransactionsRepository {
	out := AggregatedSignedTransactionsRepository{
		metaDataRepository:          metaDataRepository,
		signedTrsRepository:         signedTrsRepository,
		aggrSignedTrsBuilderFactory: aggrSignedTrsBuilderFactory,
	}

	return &out
}

// Retrieve retrieves an AggregatedSignedTransactions instance
func (rep *AggregatedSignedTransactionsRepository) Retrieve(dirPath string) (aggregated.AggregatedSignedTransactions, error) {
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
	out, outErr := rep.aggrSignedTrsBuilderFactory.Create().Create().WithMetaData(met).WithTransactions(reOrderedTrs).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
