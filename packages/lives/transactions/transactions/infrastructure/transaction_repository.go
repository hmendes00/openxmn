package infrastructure

import (
	"errors"
	"fmt"

	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	trs "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain"
)

// TransactionRepository represents a concrete TransactionRepository implementation
type TransactionRepository struct {
	objRepository objects.ObjectRepository
}

// CreateTransactionRepository creates a new TransactionRepository instance
func CreateTransactionRepository(objRepository objects.ObjectRepository) trs.TransactionRepository {
	out := TransactionRepository{
		objRepository: objRepository,
	}
	return &out
}

// Retrieve retrieves a Transaction instance
func (rep *TransactionRepository) Retrieve(dirPath string) (trs.Transaction, error) {
	obj, objErr := rep.objRepository.Retrieve(dirPath)
	if objErr != nil {
		return nil, objErr
	}

	return rep.fromObjectToTransaction(obj)
}

// RetrieveAll retrieves a []Transaction instances
func (rep *TransactionRepository) RetrieveAll(dirPath string) ([]trs.Transaction, error) {
	objs, objsErr := rep.objRepository.RetrieveAll(dirPath)
	if objsErr != nil {
		return nil, objsErr
	}

	return rep.fromObjectsToTransactions(objs)
}

func (rep *TransactionRepository) fromObjectsToTransactions(objs []objects.Object) ([]trs.Transaction, error) {
	out := []trs.Transaction{}
	for _, oneObj := range objs {
		oneTrs, oneTrsErr := rep.fromObjectToTransaction(oneObj)
		if oneTrsErr != nil {
			return nil, oneTrsErr
		}

		out = append(out, oneTrs)
	}

	return out, nil
}

func (rep *TransactionRepository) fromObjectToTransaction(obj objects.Object) (trs.Transaction, error) {
	if obj.HasSignature() {
		str := fmt.Sprintf("the transaction (id: %s) must not contain a signature", obj.GetID().String())
		return nil, errors.New(str)
	}

	if !obj.HasChunks() {
		str := fmt.Sprintf("the transaction (id: %s) must contain chunks", obj.GetID().String())
		return nil, errors.New(str)
	}

	chks := obj.GetChunks()
	newTrs := new(Transaction)
	marErr := chks.Marshal(newTrs)
	if marErr != nil {
		return nil, marErr
	}

	return newTrs, nil
}
