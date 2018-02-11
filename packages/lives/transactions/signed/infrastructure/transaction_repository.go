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
func CreateTransactionRepository(objRepository objects.ObjectRepository, transactionRepository trs.TransactionRepository, signedTrsBuilderFactory signed_trs.TransactionBuilderFactory) signed_trs.TransactionRepository {
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

	return rep.fromObjectToTransaction(obj)
}

// RetrieveAll retrieves a []Transaction instances
func (rep *TransactionRepository) RetrieveAll(dirPath string) ([]signed_trs.Transaction, error) {
	objs, objsErr := rep.objRepository.RetrieveAll(dirPath)
	if objsErr != nil {
		return nil, objsErr
	}

	return rep.fromObjectsToTransactions(objs)
}

func (rep *TransactionRepository) fromObjectsToTransactions(objs []objects.Object) ([]signed_trs.Transaction, error) {
	out := []signed_trs.Transaction{}
	for _, oneObj := range objs {
		oneTrs, oneTrsErr := rep.fromObjectToTransaction(oneObj)
		if oneTrsErr != nil {
			return nil, oneTrsErr
		}

		out = append(out, oneTrs)
	}

	return out, nil
}

func (rep *TransactionRepository) fromObjectToTransaction(obj objects.Object) (signed_trs.Transaction, error) {
	if !obj.HasSignature() {
		str := fmt.Sprintf("the signed transaction (object id: %s) must contain a signature", obj.GetID().String())
		return nil, errors.New(str)
	}

	if obj.HasChunks() {
		str := fmt.Sprintf("the signed transaction (object id: %s) must not contains chunks", obj.GetID().String())
		return nil, errors.New(str)
	}

	//retrieve the transaction:
	trsDirPath := filepath.Join(obj.GetPath(), "transaction")
	trs, trsErr := rep.transactionRepository.Retrieve(trsDirPath)
	if trsErr != nil {
		str := fmt.Sprintf("the signed transaction (id: %s) must contain a transaction in directory: %s", obj.GetID().String(), trsDirPath)
		return nil, errors.New(str)
	}

	id := obj.GetID()
	sig := obj.GetSignature()
	createdOn := obj.CreatedOn()

	//build the signed transaction:
	signedTrs, signedTrsErr := rep.signedTrsBuilderFactory.Create().Create().WithID(id).WithSignature(sig).WithTransaction(trs).CreatedOn(createdOn).Now()
	if signedTrsErr != nil {
		return nil, signedTrsErr
	}

	return signedTrs, nil
}
