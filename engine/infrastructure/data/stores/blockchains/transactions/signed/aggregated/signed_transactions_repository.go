package aggregated

import (
	"io/ioutil"
	"path/filepath"

	stored_files "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/files"
	stored_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/blockchains/transactions/signed/aggregated"
	stored_users "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/blockchains/users"
)

// SignedTransactionsRepository represents a concrete aggregated signed transactions repository implementation
type SignedTransactionsRepository struct {
	fileRepository            stored_files.FileRepository
	sigRepository             stored_users.SignatureRepository
	transRepository           stored_aggregated_transactions.TransactionsRepository
	signedTransBuilderFactory stored_aggregated_transactions.SignedTransactionsBuilderFactory
}

// CreateSignedTransactionsRepository creates a new SignedTransactionsRepository instance
func CreateSignedTransactionsRepository(fileRepository stored_files.FileRepository, sigRepository stored_users.SignatureRepository, transRepository stored_aggregated_transactions.TransactionsRepository, signedTransBuilderFactory stored_aggregated_transactions.SignedTransactionsBuilderFactory) stored_aggregated_transactions.SignedTransactionsRepository {
	out := SignedTransactionsRepository{
		fileRepository:            fileRepository,
		sigRepository:             sigRepository,
		transRepository:           transRepository,
		signedTransBuilderFactory: signedTransBuilderFactory,
	}

	return &out
}

// Retrieve retrieves an aggregated SignedTransactions instance
func (rep *SignedTransactionsRepository) Retrieve(dirPath string) (stored_aggregated_transactions.SignedTransactions, error) {
	metPath := filepath.Join(dirPath, "metadata.json")
	met, metErr := rep.fileRepository.Retrieve(metPath)
	if metErr != nil {
		return nil, metErr
	}

	sigPath := filepath.Join(dirPath, "signature")
	sig, sigErr := rep.sigRepository.Retrieve(sigPath)
	if sigErr != nil {
		return nil, sigErr
	}

	trsPath := filepath.Join(dirPath, "transactions")
	aggrTrans, aggrTransErr := rep.transRepository.Retrieve(trsPath)
	if aggrTransErr != nil {
		return nil, aggrTransErr
	}

	out, outErr := rep.signedTransBuilderFactory.Create().Create().WithMetaData(met).WithSignature(sig).WithTransactions(aggrTrans).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}

// RetrieveAll retrieves aggregated []SignedTransactions instance
func (rep *SignedTransactionsRepository) RetrieveAll(dirPath string) ([]stored_aggregated_transactions.SignedTransactions, error) {
	files, filesErr := ioutil.ReadDir(dirPath)
	if filesErr != nil {
		return nil, filesErr
	}

	out := []stored_aggregated_transactions.SignedTransactions{}
	for _, oneFile := range files {
		if !oneFile.IsDir() {
			continue
		}

		oneDirPath := filepath.Join(dirPath, oneFile.Name())
		oneTrs, oneTrsErr := rep.Retrieve(oneDirPath)
		if oneTrsErr != nil {
			return nil, oneTrsErr
		}

		out = append(out, oneTrs)
	}

	return out, nil
}
