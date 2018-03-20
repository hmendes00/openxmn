package signed

import (
	"path/filepath"

	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/files"
	stored_signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/transactions/signed"
)

// TransactionsRepository represents a stored signed transactions repository implementation
type TransactionsRepository struct {
	fileRepository            stored_files.FileRepository
	signedTrsRepository       stored_signed_transactions.TransactionRepository
	signedTransBuilderFactory stored_signed_transactions.TransactionsBuilderFactory
}

// CreateTransactionsRepository creates a new TransactionsRepository instance
func CreateTransactionsRepository(fileRepository stored_files.FileRepository, signedTrsRepository stored_signed_transactions.TransactionRepository, signedTransBuilderFactory stored_signed_transactions.TransactionsBuilderFactory) stored_signed_transactions.TransactionsRepository {
	out := TransactionsRepository{
		fileRepository:            fileRepository,
		signedTrsRepository:       signedTrsRepository,
		signedTransBuilderFactory: signedTransBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a transactions instance
func (rep *TransactionsRepository) Retrieve(dirPath string) (stored_signed_transactions.Transactions, error) {
	metPath := filepath.Join(dirPath, "metadata.json")
	met, metErr := rep.fileRepository.Retrieve(metPath)
	if metErr != nil {
		return nil, metErr
	}

	trsPath := filepath.Join(dirPath, "signed_transactions")
	signedTrs, signedTrsErr := rep.signedTrsRepository.RetrieveAll(trsPath)
	if signedTrsErr != nil {
		return nil, signedTrsErr
	}

	out, outErr := rep.signedTransBuilderFactory.Create().Create().WithMetaData(met).WithTransactions(signedTrs).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
