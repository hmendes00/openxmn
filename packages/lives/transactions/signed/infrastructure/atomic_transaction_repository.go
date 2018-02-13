package infrastructure

import (
	"errors"
	"fmt"
	"path/filepath"

	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	signed_trs "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	trs "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain"
)

// AtomicTransactionRepository represents a concrete AtomicTransactionRepository implementation
type AtomicTransactionRepository struct {
	objRepository           objects.ObjectRepository
	transactionRepository   trs.TransactionRepository
	signedTrsBuilderFactory signed_trs.TransactionBuilderFactory
	atomicTrsBuilderFactory signed_trs.AtomicTransactionBuilderFactory
}

// CreateAtomicTransactionRepository creates a new AtomicTransactionRepository instance
func CreateAtomicTransactionRepository(objRepository objects.ObjectRepository, transactionRepository trs.TransactionRepository, signedTrsBuilderFactory signed_trs.TransactionBuilderFactory, atomicTrsBuilderFactory signed_trs.AtomicTransactionBuilderFactory) signed_trs.AtomicTransactionRepository {
	out := AtomicTransactionRepository{
		objRepository:           objRepository,
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

	return rep.fromObjectToAtomicTransaction(obj)
}

// RetrieveAll retrieves a []AtomicTransaction instances
func (rep *AtomicTransactionRepository) RetrieveAll(dirPath string) ([]signed_trs.AtomicTransaction, error) {
	objs, objsErr := rep.objRepository.RetrieveAll(dirPath)
	if objsErr != nil {
		return nil, objsErr
	}

	return rep.fromObjectsToAtomicTransactions(objs)
}

func (rep *AtomicTransactionRepository) fromObjectsToAtomicTransactions(objs []objects.Object) ([]signed_trs.AtomicTransaction, error) {
	out := []signed_trs.AtomicTransaction{}
	for _, oneObj := range objs {
		oneTrs, oneTrsErr := rep.fromObjectToAtomicTransaction(oneObj)
		if oneTrsErr != nil {
			return nil, oneTrsErr
		}

		out = append(out, oneTrs)
	}

	return out, nil
}

func (rep *AtomicTransactionRepository) fromObjectToAtomicTransaction(obj objects.Object) (signed_trs.AtomicTransaction, error) {
	if !obj.HasSignature() {
		str := fmt.Sprintf("the signed transaction (id: %s) must contain a signature", obj.GetID().String())
		return nil, errors.New(str)
	}

	//retrieve the transactions:
	trsDirPath := filepath.Join("transactions")
	trs, trsErr := rep.transactionRepository.RetrieveAll(trsDirPath)
	if trsErr != nil {
		str := fmt.Sprintf("the signed transactions (object id: %s) must contain []AtomicTransaction in directory: %s", obj.GetID().String(), trsDirPath)
		return nil, errors.New(str)
	}

	id := obj.GetID()
	sig := obj.GetSignature()
	createdOn := obj.CreatedOn()

	//build the atomic transaction:
	atomicTrs, atomicTrsErr := rep.atomicTrsBuilderFactory.Create().Create().WithID(id).WithSignature(sig).WithTransactions(trs).CreatedOn(createdOn).Now()
	if atomicTrsErr != nil {
		return nil, atomicTrsErr
	}

	return atomicTrs, nil
}
