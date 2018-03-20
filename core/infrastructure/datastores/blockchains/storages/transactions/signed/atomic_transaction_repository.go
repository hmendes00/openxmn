package signed

import (
	"io/ioutil"
	"path/filepath"

	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/files"
	stored_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/transactions"
	stored_signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/transactions/signed"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/users"
)

// AtomicTransactionRepository represents a concrete stored AtomicTransactionRepository implementation
type AtomicTransactionRepository struct {
	fileRepository          stored_files.FileRepository
	sigRepository           stored_users.SignatureRepository
	transRepository         stored_transactions.Repository
	atomicTrsBuilderFactory stored_signed_transactions.AtomicTransactionBuilderFactory
}

// CreateAtomicTransactionRepository creates a new AtomicTransactionRepository instance
func CreateAtomicTransactionRepository(fileRepository stored_files.FileRepository, sigRepository stored_users.SignatureRepository, transRepository stored_transactions.Repository, atomicTrsBuilderFactory stored_signed_transactions.AtomicTransactionBuilderFactory) stored_signed_transactions.AtomicTransactionRepository {
	out := AtomicTransactionRepository{
		fileRepository:          fileRepository,
		sigRepository:           sigRepository,
		transRepository:         transRepository,
		atomicTrsBuilderFactory: atomicTrsBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a stored atomic transaction
func (rep *AtomicTransactionRepository) Retrieve(dirPath string) (stored_signed_transactions.AtomicTransaction, error) {
	metPath := filepath.Join(dirPath, "metadata.json")
	met, metErr := rep.fileRepository.Retrieve(metPath)
	if metErr != nil {
		return nil, metErr
	}

	trsDirPath := filepath.Join(dirPath, "transactions")
	trs, trsErr := rep.transRepository.Retrieve(trsDirPath)
	if trsErr != nil {
		return nil, trsErr
	}

	sigPath := filepath.Join(dirPath, "signature")
	sig, sigErr := rep.sigRepository.Retrieve(sigPath)
	if sigErr != nil {
		return nil, sigErr
	}

	out, outErr := rep.atomicTrsBuilderFactory.Create().Create().WithMetaData(met).WithSignature(sig).WithTransactions(trs).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}

// RetrieveAll retrieves a stored []AtomicTransaction
func (rep *AtomicTransactionRepository) RetrieveAll(dirPath string) ([]stored_signed_transactions.AtomicTransaction, error) {
	files, filesErr := ioutil.ReadDir(dirPath)
	if filesErr != nil {
		return nil, filesErr
	}

	out := []stored_signed_transactions.AtomicTransaction{}
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
