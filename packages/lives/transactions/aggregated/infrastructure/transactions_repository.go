package infrastructure

import (
	"encoding/hex"
	"errors"
	"fmt"
	"path/filepath"

	hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/domain"
	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/domain"
	signed "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
)

// TransactionsRepository represents a concrete TransactionsRepository implementation
type TransactionsRepository struct {
	signedTrsRepository         signed.TransactionRepository
	atomicSignedTrsRepository   signed.AtomicTransactionRepository
	htRepository                hashtrees.HashTreeRepository
	objectRepository            objects.ObjectRepository
	aggregatedTrsBuilderFactory aggregated.TransactionsBuilderFactory
}

// CreateTransactionsRepository creates a new TransactionsRepository instance
func CreateTransactionsRepository(
	signedTrsRepository signed.TransactionRepository,
	atomicSignedTrsRepository signed.AtomicTransactionRepository,
	htRepository hashtrees.HashTreeRepository,
	objectRepository objects.ObjectRepository,
	aggregatedTrsBuilderFactory aggregated.TransactionsBuilderFactory,
) aggregated.TransactionsRepository {
	out := TransactionsRepository{
		signedTrsRepository:         signedTrsRepository,
		atomicSignedTrsRepository:   atomicSignedTrsRepository,
		htRepository:                htRepository,
		objectRepository:            objectRepository,
		aggregatedTrsBuilderFactory: aggregatedTrsBuilderFactory,
	}
	return &out
}

// Retrieve retrieves an aggregated Transactions instance
func (rep *TransactionsRepository) Retrieve(dirPath string) (aggregated.Transactions, error) {

	//retrieve the hashtree:
	ht, htErr := rep.htRepository.Retrieve(dirPath)
	if htErr != nil {
		return nil, htErr
	}

	//retrieve the obj:
	obj, objErr := rep.objectRepository.Retrieve(dirPath)
	if objErr != nil {
		return nil, objErr
	}

	if obj.HasChunks() {
		return nil, errors.New("there should be no chunks inside the Transactions instance")
	}

	//declare the ht blocks:
	htBlocks := [][]byte{}

	//retrieve the signed transaction, if any - and build the hashtree blocks:
	signedTrsMap := map[string]signed.Transaction{}
	signedTrsPath := filepath.Join(dirPath, "signed_transactions")
	signedTrs, signedTrsErr := rep.signedTrsRepository.RetrieveAll(signedTrsPath)
	if signedTrsErr == nil {
		for _, oneSignedTrs := range signedTrs {
			idAsBytes := oneSignedTrs.GetID().Bytes()
			idAsString := hex.EncodeToString(idAsBytes)
			signedTrsMap[idAsString] = oneSignedTrs
			htBlocks = append(htBlocks, idAsBytes)
		}
	}

	//retrieve the atomic transaction, if any - and build the hashtree blocks:
	atomicTrsMap := map[string]signed.AtomicTransaction{}
	atomicTrsPath := filepath.Join(dirPath, "signed_atomic_transactions")
	atomicTrs, atomicTrsErr := rep.atomicSignedTrsRepository.RetrieveAll(atomicTrsPath)
	if atomicTrsErr == nil {
		for _, oneAtomicTrs := range atomicTrs {
			idAsBytes := oneAtomicTrs.GetID().Bytes()
			idAsString := hex.EncodeToString(idAsBytes)
			atomicTrsMap[idAsString] = oneAtomicTrs
			htBlocks = append(htBlocks, idAsBytes)
		}
	}

	if signedTrsErr != nil && atomicTrsErr != nil {
		str := fmt.Sprintf("There was no signed transactions (error: %s) or signed atomic transactions (error: %s)", signedTrsErr.Error(), atomicTrsErr.Error())
		return nil, errors.New(str)
	}

	//re-order the block:
	reOrderedBlocks, reOrderedBlocksErr := ht.Order(htBlocks)
	if reOrderedBlocksErr != nil {
		return nil, reOrderedBlocksErr
	}

	//build the transactions instance:
	amountTrs := 0
	met := obj.GetMetaData()
	id := met.GetID()
	createdOn := met.CreatedOn()
	trsBuilder := rep.aggregatedTrsBuilderFactory.Create().Create().WithID(id).CreatedOn(createdOn)
	if signedTrsErr == nil {
		//re-order the signed transactions:
		orderedTrs := []signed.Transaction{}
		for _, oneOrderedBlk := range reOrderedBlocks {
			idAsString := hex.EncodeToString(oneOrderedBlk)
			if oneSignedTrs, ok := signedTrsMap[idAsString]; ok {
				orderedTrs = append(orderedTrs, oneSignedTrs)
				amountTrs++
				continue
			}
		}

		trsBuilder.WithTransactions(orderedTrs)
	}

	//build the atomic transactions instance:
	if atomicTrsErr == nil {
		//re-order the signed atomic transactions:
		orderedAtomicTrs := []signed.AtomicTransaction{}
		for _, oneOrderedBlk := range reOrderedBlocks {
			idAsString := hex.EncodeToString(oneOrderedBlk)
			if oneAtomicTrs, ok := atomicTrsMap[idAsString]; ok {
				orderedAtomicTrs = append(orderedAtomicTrs, oneAtomicTrs)
				amountTrs++
				continue
			}
		}

		trsBuilder.WithAtomicTransactions(orderedAtomicTrs)
	}

	if amountTrs != len(reOrderedBlocks) {
		str := fmt.Sprintf("the amount of transactions (%d) does not match the amount or hashtree blocks (%d)", amountTrs, len(reOrderedBlocks))
		return nil, errors.New(str)
	}

	trs, trsErr := trsBuilder.Now()
	if trsErr != nil {
		return nil, trsErr
	}

	return trs, nil
}
