package infrastructure

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/domain"
)

// SignedTransactionsRepository represents a concrete SignedTransactions repository
type SignedTransactionsRepository struct {
	objRepository           objects.ObjectRepository
	trsRepository           aggregated.TransactionsRepository
	signedTrsBuilderFactory aggregated.SignedTransactionsBuilderFactory
}

// CreateSignedTransactionsRepository creates a new SignedTransactionsRepository instance
func CreateSignedTransactionsRepository(
	objRepository objects.ObjectRepository,
	trsRepository aggregated.TransactionsRepository,
	signedTrsBuilderFactory aggregated.SignedTransactionsBuilderFactory,
) aggregated.SignedTransactionsRepository {
	out := SignedTransactionsRepository{
		objRepository:           objRepository,
		trsRepository:           trsRepository,
		signedTrsBuilderFactory: signedTrsBuilderFactory,
	}
	return &out
}

// Retrieve retrieves a SignedTransactions instance
func (rep *SignedTransactionsRepository) Retrieve(dirPath string) (aggregated.SignedTransactions, error) {
	//retrieve the object:
	obj, objErr := rep.objRepository.Retrieve(dirPath)
	if objErr != nil {
		return nil, objErr
	}

	//retrieve the transactions:
	trsPath := filepath.Join(dirPath, "transactions")
	trs, trsErr := rep.trsRepository.Retrieve(trsPath)
	if trsErr != nil {
		return nil, trsErr
	}

	met := obj.GetMetaData()
	if !met.HasSignature() {
		return nil, errors.New("the signature, inside the metadata is mandatory")
	}

	//build the signed transactions:
	id := met.GetID()
	ts := met.CreatedOn()
	sig := met.GetSignature()
	signedTrs, signedTrsErr := rep.signedTrsBuilderFactory.Create().Create().WithID(id).CreatedOn(ts).WithSignature(sig).WithTransactions(trs).Now()
	if signedTrsErr != nil {
		return nil, signedTrsErr
	}

	//return:
	return signedTrs, nil
}

// RetrieveAll retrieves []SignedTransactions instances
func (rep *SignedTransactionsRepository) RetrieveAll(dirPath string) ([]aggregated.SignedTransactions, error) {
	files, filesErr := ioutil.ReadDir(dirPath)
	if filesErr != nil {
		return nil, filesErr
	}

	signedTrs := []aggregated.SignedTransactions{}
	for _, oneFile := range files {
		if !oneFile.IsDir() {
			continue
		}

		trsDirPath := filepath.Join(dirPath, oneFile.Name())
		oneSignedTrs, oneSignedTrsErr := rep.Retrieve(trsDirPath)
		if oneSignedTrsErr != nil {
			return nil, oneSignedTrsErr
		}

		signedTrs = append(signedTrs, oneSignedTrs)
	}

	return signedTrs, nil
}
