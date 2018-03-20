package transactions

import (
	"io/ioutil"
	"path/filepath"

	stored_chunks "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/chunks"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/files"
	stored_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/transactions"
)

// TransactionRepository represents a concrete stored transaction repository implementation
type TransactionRepository struct {
	fileRepository    stored_files.FileRepository
	chkRepository     stored_chunks.Repository
	trsBuilderFactory stored_transactions.TransactionBuilderFactory
}

// CreateTransactionRepository creates a new TransactionRepository instance
func CreateTransactionRepository(fileRepository stored_files.FileRepository, chkRepository stored_chunks.Repository, trsBuilderFactory stored_transactions.TransactionBuilderFactory) stored_transactions.TransactionRepository {
	out := TransactionRepository{
		fileRepository:    fileRepository,
		chkRepository:     chkRepository,
		trsBuilderFactory: trsBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a stored Transaction
func (rep *TransactionRepository) Retrieve(dirPath string) (stored_transactions.Transaction, error) {
	metDirPath := filepath.Join(dirPath, "metadata.json")
	met, metErr := rep.fileRepository.Retrieve(metDirPath)
	if metErr != nil {
		return nil, metErr
	}

	chkDirPath := filepath.Join(dirPath, "json")
	chk, chkErr := rep.chkRepository.Retrieve(chkDirPath)
	if chkErr != nil {
		return nil, chkErr
	}

	out, outErr := rep.trsBuilderFactory.Create().Create().WithMetaData(met).WithChunks(chk).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}

// RetrieveAll retrieves a stored []Transaction
func (rep *TransactionRepository) RetrieveAll(dirPath string) ([]stored_transactions.Transaction, error) {
	files, filesErr := ioutil.ReadDir(dirPath)
	if filesErr != nil {
		return nil, filesErr
	}

	out := []stored_transactions.Transaction{}
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
