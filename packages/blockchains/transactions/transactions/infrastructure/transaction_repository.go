package infrastructure

import (
	"io/ioutil"
	"path/filepath"

	chunks "github.com/XMNBlockchain/core/packages/blockchains/chunks/domain"
	trs "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
)

// TransactionRepository represents a concrete TransactionRepository implementation
type TransactionRepository struct {
	chksRepository chunks.ChunksRepository
}

// CreateTransactionRepository creates a new TransactionRepository instance
func CreateTransactionRepository(chksRepository chunks.ChunksRepository) trs.TransactionRepository {
	out := TransactionRepository{
		chksRepository: chksRepository,
	}
	return &out
}

// Retrieve retrieves a Transaction instance
func (rep *TransactionRepository) Retrieve(dirPath string) (trs.Transaction, error) {
	//retrieve the chunks:
	chks, chksErr := rep.chksRepository.Retrieve(dirPath)
	if chksErr != nil {
		return nil, chksErr
	}

	newTrs := new(Transaction)
	marErr := chks.Marshal(newTrs)
	if marErr != nil {
		return nil, marErr
	}

	return newTrs, nil
}

// RetrieveAll retrieves a []Transaction instances
func (rep *TransactionRepository) RetrieveAll(dirPath string) ([]trs.Transaction, error) {
	files, filesErr := ioutil.ReadDir(dirPath)
	if filesErr != nil {
		return nil, filesErr
	}

	out := []trs.Transaction{}
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
