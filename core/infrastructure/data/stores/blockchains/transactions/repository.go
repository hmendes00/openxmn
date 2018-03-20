package transactions

import (
	"path/filepath"

	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/files"
	stored_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/transactions"
)

// Repository represents a concrete stored transactions repository implementation
type Repository struct {
	fileRepository      stored_files.FileRepository
	trsRepository       stored_transactions.TransactionRepository
	transBuilderFactory stored_transactions.BuilderFactory
}

// CreateRepository creates a new repository instance
func CreateRepository(fileRepository stored_files.FileRepository, trsRepository stored_transactions.TransactionRepository, transBuilderFactory stored_transactions.BuilderFactory) stored_transactions.Repository {
	out := Repository{
		fileRepository:      fileRepository,
		trsRepository:       trsRepository,
		transBuilderFactory: transBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a Transactions instance
func (rep *Repository) Retrieve(dirPath string) (stored_transactions.Transactions, error) {
	metDirPath := filepath.Join(dirPath, "metadata.json")
	met, metErr := rep.fileRepository.Retrieve(metDirPath)
	if metErr != nil {
		return nil, metErr
	}

	trsDirPath := filepath.Join(dirPath, "transactions")
	trs, trsErr := rep.trsRepository.RetrieveAll(trsDirPath)
	if trsErr != nil {
		return nil, trsErr
	}

	out, outErr := rep.transBuilderFactory.Create().Create().WithMetaData(met).WithTransactions(trs).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
