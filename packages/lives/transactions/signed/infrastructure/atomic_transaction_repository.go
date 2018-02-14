package infrastructure

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/infrastructure"
	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	signed_trs "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	transactions "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain"
)

// AtomicTransactionRepository represents a concrete AtomicTransactionRepository implementation
type AtomicTransactionRepository struct {
	objRepository           objects.ObjectRepository
	fileRepository          files.FileRepository
	transactionRepository   transactions.TransactionRepository
	signedTrsBuilderFactory signed_trs.TransactionBuilderFactory
	atomicTrsBuilderFactory signed_trs.AtomicTransactionBuilderFactory
}

// CreateAtomicTransactionRepository creates a new AtomicTransactionRepository instance
func CreateAtomicTransactionRepository(
	objRepository objects.ObjectRepository,
	fileRepository files.FileRepository,
	transactionRepository transactions.TransactionRepository,
	signedTrsBuilderFactory signed_trs.TransactionBuilderFactory,
	atomicTrsBuilderFactory signed_trs.AtomicTransactionBuilderFactory,
) signed_trs.AtomicTransactionRepository {
	out := AtomicTransactionRepository{
		objRepository:           objRepository,
		fileRepository:          fileRepository,
		transactionRepository:   transactionRepository,
		signedTrsBuilderFactory: signedTrsBuilderFactory,
		atomicTrsBuilderFactory: atomicTrsBuilderFactory,
	}
	return &out
}

// Retrieve retrieves a AtomicTransaction instance
func (rep *AtomicTransactionRepository) Retrieve(dirPath string) (signed_trs.AtomicTransaction, error) {
	obj, objErr := rep.objRepository.Retrieve(dirPath)
	if objErr != nil {
		return nil, objErr
	}

	return rep.fromObjectToAtomicTransaction(dirPath, obj)
}

// RetrieveAll retrieves a []AtomicTransaction instances
func (rep *AtomicTransactionRepository) RetrieveAll(dirPath string) ([]signed_trs.AtomicTransaction, error) {
	objs, objsErr := rep.objRepository.RetrieveAll(dirPath)
	if objsErr != nil {
		return nil, objsErr
	}

	return rep.fromObjectsToAtomicTransactions(dirPath, objs)
}

func (rep *AtomicTransactionRepository) fromObjectsToAtomicTransactions(dirPath string, objs []objects.Object) ([]signed_trs.AtomicTransaction, error) {
	out := []signed_trs.AtomicTransaction{}
	for _, oneObj := range objs {
		oneAtomicTrsDirPath := filepath.Join(dirPath, oneObj.GetMetaData().GetID().String())
		oneTrs, oneTrsErr := rep.fromObjectToAtomicTransaction(oneAtomicTrsDirPath, oneObj)
		if oneTrsErr != nil {
			return nil, oneTrsErr
		}

		out = append(out, oneTrs)
	}

	return out, nil
}

func (rep *AtomicTransactionRepository) fromObjectToAtomicTransaction(dirPath string, obj objects.Object) (signed_trs.AtomicTransaction, error) {

	if obj.HasChunks() {
		str := fmt.Sprintf("the atomic transaction (object id: %s) must not contain chunks", obj.GetMetaData().GetID().String())
		return nil, errors.New(str)
	}

	//retrieve the transactions:
	trsDirPath := filepath.Join(dirPath, "transactions")
	trs, trsErr := rep.transactionRepository.RetrieveAll(trsDirPath)
	if trsErr != nil {
		str := fmt.Sprintf("the atomic transactions (object id: %s) must contain []AtomicTransaction in directory: %s", obj.GetMetaData().GetID().String(), trsDirPath)
		return nil, errors.New(str)
	}

	trsMap := map[string]transactions.Transaction{}
	for _, oneTrs := range trs {
		trsIDAsString := hex.EncodeToString(oneTrs.GetID().Bytes())
		trsMap[trsIDAsString] = oneTrs
	}

	//read the hashtree:
	htFile, htFileErr := rep.fileRepository.Retrieve(dirPath, "hashtree.json")
	if htFileErr != nil {
		return nil, htFileErr
	}

	//unmarshal the hashtree:
	newHt := new(concrete_hashtrees.HashTree)
	jsonErr := json.Unmarshal(htFile.GetData(), newHt)
	if jsonErr != nil {
		return nil, jsonErr
	}

	//re-order the transactions:
	trsIDs := [][]byte{}
	for _, oneTrs := range trs {
		trsIDs = append(trsIDs, oneTrs.GetID().Bytes())
	}

	orderedData, orderedErr := newHt.Order(trsIDs)
	if orderedErr != nil {
		return nil, orderedErr
	}

	//create the new trs list based on the ordered data:
	orderedTrs := []transactions.Transaction{}
	for _, oneOrderedData := range orderedData {
		trsIDAsString := hex.EncodeToString(oneOrderedData)
		if oneTrs, ok := trsMap[trsIDAsString]; ok {
			orderedTrs = append(orderedTrs, oneTrs)
			continue
		}

		str := fmt.Sprintf("the ordered transaction ID (%s), in the HashTree, cannot be found.  Path: %s. \n\n%v", trsIDAsString, dirPath, trsMap)
		return nil, errors.New(str)
	}

	metaData := obj.GetMetaData()
	id := metaData.GetID()
	sig := metaData.GetSignature()
	createdOn := metaData.CreatedOn()

	//build the atomic transaction:
	atomicTrs, atomicTrsErr := rep.atomicTrsBuilderFactory.Create().Create().WithID(id).WithSignature(sig).WithTransactions(orderedTrs).CreatedOn(createdOn).Now()
	if atomicTrsErr != nil {
		return nil, atomicTrsErr
	}

	return atomicTrs, nil
}
