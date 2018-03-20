package aggregated

import (
	"path/filepath"

	stored_files "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/files"
	stored_signed_transactions "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/blockchains/transactions/signed"
	stored_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/blockchains/transactions/signed/aggregated"
)

// TransactionsRepository represents a concrete aggregated transactions repository implementation
type TransactionsRepository struct {
	fileRepository          stored_files.FileRepository
	transRepository         stored_signed_transactions.TransactionsRepository
	atomicTransRepository   stored_signed_transactions.AtomicTransactionsRepository
	aggrTransBuilderFactory stored_aggregated_transactions.TransactionsBuilderFactory
}

// CreateTransactionsRepository creates a new TransactionsRepository instance
func CreateTransactionsRepository(fileRepository stored_files.FileRepository, transRepository stored_signed_transactions.TransactionsRepository, atomicTransRepository stored_signed_transactions.AtomicTransactionsRepository, aggrTransBuilderFactory stored_aggregated_transactions.TransactionsBuilderFactory) stored_aggregated_transactions.TransactionsRepository {
	out := TransactionsRepository{
		fileRepository:          fileRepository,
		transRepository:         transRepository,
		atomicTransRepository:   atomicTransRepository,
		aggrTransBuilderFactory: aggrTransBuilderFactory,
	}

	return &out
}

// Retrieve retrieves an aggregated Transactions instance
func (rep *TransactionsRepository) Retrieve(dirPath string) (stored_aggregated_transactions.Transactions, error) {
	metPath := filepath.Join(dirPath, "metadata.json")
	met, metErr := rep.fileRepository.Retrieve(metPath)
	if metErr != nil {
		return nil, metErr
	}

	//reate the aggregated transactions builder:
	aggrTransBuilder := rep.aggrTransBuilderFactory.Create().Create().WithMetaData(met)

	//if there is transactions:
	trsPath := filepath.Join(dirPath, "signed_transactions")
	trans, transErr := rep.transRepository.Retrieve(trsPath)
	if transErr == nil {
		aggrTransBuilder.WithTransactions(trans)
	}

	//if there is atomic transactions:
	atomicTrsPath := filepath.Join(dirPath, "signed_atomic_transactions")
	atomicTrans, atomicTransErr := rep.atomicTransRepository.Retrieve(atomicTrsPath)
	if atomicTransErr == nil {
		aggrTransBuilder.WithAtomicTransactions(atomicTrans)
	}

	out, outErr := aggrTransBuilder.Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
