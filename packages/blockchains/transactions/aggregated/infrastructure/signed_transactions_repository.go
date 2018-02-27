package infrastructure

import (
	"io/ioutil"
	"path/filepath"

	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
)

// SignedTransactionsRepository represents a concrete SignedTransactions repository
type SignedTransactionsRepository struct {
	metaDataRepository      metadata.MetaDataRepository
	userSigRepository       users.SignatureRepository
	aggregatedTrsRepository aggregated.TransactionsRepository
	signedTrsBuilderFactory aggregated.SignedTransactionsBuilderFactory
}

// CreateSignedTransactionsRepository creates a new SignedTransactionsRepository instance
func CreateSignedTransactionsRepository(
	metaDataRepository metadata.MetaDataRepository,
	userSigRepository users.SignatureRepository,
	aggregatedTrsRepository aggregated.TransactionsRepository,
	signedTrsBuilderFactory aggregated.SignedTransactionsBuilderFactory,
) aggregated.SignedTransactionsRepository {
	out := SignedTransactionsRepository{
		metaDataRepository:      metaDataRepository,
		userSigRepository:       userSigRepository,
		aggregatedTrsRepository: aggregatedTrsRepository,
		signedTrsBuilderFactory: signedTrsBuilderFactory,
	}
	return &out
}

// Retrieve retrieves a SignedTransactions instance
func (rep *SignedTransactionsRepository) Retrieve(dirPath string) (aggregated.SignedTransactions, error) {
	//retrieve the metadata:
	met, metErr := rep.metaDataRepository.Retrieve(dirPath)
	if metErr != nil {
		return nil, metErr
	}

	//retrieve the signature:
	sig, sigErr := rep.userSigRepository.Retrieve(dirPath)
	if sigErr != nil {
		return nil, sigErr
	}

	//retrieve the transactions:
	trsPath := filepath.Join(dirPath, "transactions")
	trs, trsErr := rep.aggregatedTrsRepository.Retrieve(trsPath)
	if trsErr != nil {
		return nil, trsErr
	}

	//build the signed transactions:
	id := met.GetID()
	ts := met.CreatedOn()
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
