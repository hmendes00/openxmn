package signed

import (
	"path/filepath"

	stored_files "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/files"
	stored_signed_transactions "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/blockchains/transactions/signed"
)

// AtomicTransactionsRepository represents a stored AtomicTransactionsRepository implementation
type AtomicTransactionsRepository struct {
	fileRepository            stored_files.FileRepository
	atomicTrsRepository       stored_signed_transactions.AtomicTransactionRepository
	atomicTransBuilderFactory stored_signed_transactions.AtomicTransactionsBuilderFactory
}

// CreateAtomicTransactionsRepository creates a new AtomicTransactionsRepository instance
func CreateAtomicTransactionsRepository(fileRepository stored_files.FileRepository, atomicTrsRepository stored_signed_transactions.AtomicTransactionRepository, atomicTransBuilderFactory stored_signed_transactions.AtomicTransactionsBuilderFactory) stored_signed_transactions.AtomicTransactionsRepository {
	out := AtomicTransactionsRepository{
		fileRepository:            fileRepository,
		atomicTrsRepository:       atomicTrsRepository,
		atomicTransBuilderFactory: atomicTransBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a stored AtomicTransactions instance
func (rep *AtomicTransactionsRepository) Retrieve(dirPath string) (stored_signed_transactions.AtomicTransactions, error) {
	metPath := filepath.Join(dirPath, "metadata.json")
	met, metErr := rep.fileRepository.Retrieve(metPath)
	if metErr != nil {
		return nil, metErr
	}

	atomicTrsPath := filepath.Join(dirPath, "atomic_transactions")
	atomicTrs, atomicTrsErr := rep.atomicTrsRepository.RetrieveAll(atomicTrsPath)
	if atomicTrsErr != nil {
		return nil, atomicTrsErr
	}

	out, outErr := rep.atomicTransBuilderFactory.Create().Create().WithMetaData(met).WithTransactions(atomicTrs).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
