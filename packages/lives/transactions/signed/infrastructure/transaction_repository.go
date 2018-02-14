package infrastructure

import (
	"errors"
	"fmt"
	"path/filepath"

	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	signed_trs "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	trs "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain"
)

// TransactionRepository represents a concrete TransactionRepository implementation
type TransactionRepository struct {
	objRepository           objects.ObjectRepository
	transactionRepository   trs.TransactionRepository
	signedTrsBuilderFactory signed_trs.TransactionBuilderFactory
}

// CreateTransactionRepository creates a new TransactionRepository instance
func CreateTransactionRepository(
	objRepository objects.ObjectRepository,
	transactionRepository trs.TransactionRepository,
	signedTrsBuilderFactory signed_trs.TransactionBuilderFactory,
) signed_trs.TransactionRepository {
	out := TransactionRepository{
		objRepository:           objRepository,
		transactionRepository:   transactionRepository,
		signedTrsBuilderFactory: signedTrsBuilderFactory,
	}
	return &out
}

// Retrieve retrieves a Transaction instance
func (rep *TransactionRepository) Retrieve(dirPath string) (signed_trs.Transaction, error) {
	obj, objErr := rep.objRepository.Retrieve(dirPath)
	if objErr != nil {
		return nil, objErr
	}

	signedTrs, signedTrsErr := rep.fromObjectToTransaction(dirPath, obj)
	if signedTrsErr != nil {
		return nil, signedTrsErr
	}

	return signedTrs, nil
}

// RetrieveAll retrieves a []Transaction instances
func (rep *TransactionRepository) RetrieveAll(dirPath string) ([]signed_trs.Transaction, error) {
	objs, objsErr := rep.objRepository.RetrieveAll(dirPath)
	if objsErr != nil {
		return nil, objsErr
	}

	signedTrs, signedTrsErr := rep.fromObjectsToTransactions(dirPath, objs)
	if signedTrsErr != nil {
		return nil, signedTrsErr
	}

	return signedTrs, nil
}

func (rep *TransactionRepository) fromObjectsToTransactions(dirPath string, objs []objects.Object) ([]signed_trs.Transaction, error) {
	out := []signed_trs.Transaction{}
	for _, oneObj := range objs {
		oneTrsDirPath := filepath.Join(dirPath, oneObj.GetMetaData().GetID().String())
		oneTrs, oneTrsErr := rep.fromObjectToTransaction(oneTrsDirPath, oneObj)
		if oneTrsErr != nil {
			return nil, oneTrsErr
		}

		out = append(out, oneTrs)
	}

	return out, nil
}

func (rep *TransactionRepository) fromObjectToTransaction(dirPath string, obj objects.Object) (signed_trs.Transaction, error) {
	if obj.HasChunks() {
		str := fmt.Sprintf("the signed transaction (object id: %s) must not contain chunks", obj.GetMetaData().GetID().String())
		return nil, errors.New(str)
	}

	//retrieve the transaction:
	trsDirPath := filepath.Join(dirPath, "transaction")
	trs, trsErr := rep.transactionRepository.Retrieve(trsDirPath)
	if trsErr != nil {
		return nil, trsErr
	}

	metaData := obj.GetMetaData()
	id := metaData.GetID()
	sig := metaData.GetSignature()
	createdOn := metaData.CreatedOn()

	//build the signed transaction:
	signedTrs, signedTrsErr := rep.signedTrsBuilderFactory.Create().Create().WithID(id).WithSignature(sig).WithTransaction(trs).CreatedOn(createdOn).Now()
	if signedTrsErr != nil {
		return nil, signedTrsErr
	}

	return signedTrs, nil
}
