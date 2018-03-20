package signed

import (
	"io/ioutil"
	"path/filepath"

	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	stored_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions"
	stored_signed_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions/signed"
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/users"
)

// TransactionRepository represents a concrete stored transaction repository
type TransactionRepository struct {
	fileRepository          stored_files.FileRepository
	sigRepository           stored_users.SignatureRepository
	trsRepository           stored_transactions.TransactionRepository
	signedTrsBuilderFactory stored_signed_transactions.TransactionBuilderFactory
}

// CreateTransactionRepository creates a new TransactionRepository instance
func CreateTransactionRepository(fileRepository stored_files.FileRepository, sigRepository stored_users.SignatureRepository, trsRepository stored_transactions.TransactionRepository, signedTrsBuilderFactory stored_signed_transactions.TransactionBuilderFactory) stored_signed_transactions.TransactionRepository {
	out := TransactionRepository{
		fileRepository:          fileRepository,
		sigRepository:           sigRepository,
		trsRepository:           trsRepository,
		signedTrsBuilderFactory: signedTrsBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a stored signed Transaction
func (rep *TransactionRepository) Retrieve(dirPath string) (stored_signed_transactions.Transaction, error) {
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

	trsDirPath := filepath.Join(dirPath, "transaction")
	trs, trsErr := rep.trsRepository.Retrieve(trsDirPath)
	if trsErr != nil {
		return nil, trsErr
	}

	out, outErr := rep.signedTrsBuilderFactory.Create().Create().WithMetaData(met).WithSignature(sig).WithTransaction(trs).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}

// RetrieveAll retrieves a stored signed []Transaction
func (rep *TransactionRepository) RetrieveAll(dirPath string) ([]stored_signed_transactions.Transaction, error) {
	files, filesErr := ioutil.ReadDir(dirPath)
	if filesErr != nil {
		return nil, filesErr
	}

	out := []stored_signed_transactions.Transaction{}
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
