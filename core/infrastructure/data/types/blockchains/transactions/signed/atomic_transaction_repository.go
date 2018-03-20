package signed

import (
	"io/ioutil"
	"path/filepath"

	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/metadata"
	transactions "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions"
	signed_trs "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions/signed"
	users "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/users"
)

// AtomicTransactionRepository represents a concrete AtomicTransactionRepository implementation
type AtomicTransactionRepository struct {
	metaDataRepository      metadata.MetaDataRepository
	userSigRepository       users.SignatureRepository
	trsRepository           transactions.TransactionsRepository
	atomicTrsBuilderFactory signed_trs.AtomicTransactionBuilderFactory
}

// CreateAtomicTransactionRepository creates a new AtomicTransactionRepository instance
func CreateAtomicTransactionRepository(
	metaDataRepository metadata.MetaDataRepository,
	userSigRepository users.SignatureRepository,
	trsRepository transactions.TransactionsRepository,
	atomicTrsBuilderFactory signed_trs.AtomicTransactionBuilderFactory,
) signed_trs.AtomicTransactionRepository {
	out := AtomicTransactionRepository{
		metaDataRepository:      metaDataRepository,
		userSigRepository:       userSigRepository,
		trsRepository:           trsRepository,
		atomicTrsBuilderFactory: atomicTrsBuilderFactory,
	}
	return &out
}

// Retrieve retrieves a AtomicTransaction instance
func (rep *AtomicTransactionRepository) Retrieve(dirPath string) (signed_trs.AtomicTransaction, error) {
	//retrieve the metadata:
	met, metErr := rep.metaDataRepository.Retrieve(dirPath)
	if metErr != nil {
		return nil, metErr
	}

	//retrieve the signature:
	sigPath := filepath.Join(dirPath, "signature")
	sig, sigErr := rep.userSigRepository.Retrieve(sigPath)
	if sigErr != nil {
		return nil, sigErr
	}

	//retrieve the transactions:
	trsDirPath := filepath.Join(dirPath, "transactions")
	trs, trsErr := rep.trsRepository.Retrieve(trsDirPath)
	if trsErr != nil {
		return nil, trsErr
	}

	//build the atomic transaction:
	atomicTrs, atomicTrsErr := rep.atomicTrsBuilderFactory.Create().Create().WithMetaData(met).WithSignature(sig).WithTransactions(trs).Now()
	if atomicTrsErr != nil {
		return nil, atomicTrsErr
	}

	return atomicTrs, nil
}

// RetrieveAll retrieves a []AtomicTransaction instances
func (rep *AtomicTransactionRepository) RetrieveAll(dirPath string) ([]signed_trs.AtomicTransaction, error) {
	files, filesErr := ioutil.ReadDir(dirPath)
	if filesErr != nil {
		return nil, filesErr
	}

	out := []signed_trs.AtomicTransaction{}
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
