package signed

import (
	"io/ioutil"
	"path/filepath"

	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
	trs "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/transactions"
	signed_trs "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/transactions/signed"
	users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/users"
)

// TransactionRepository represents a concrete TransactionRepository implementation
type TransactionRepository struct {
	metaDataRepository      metadata.MetaDataRepository
	sigRepository           users.SignatureRepository
	transactionRepository   trs.TransactionRepository
	signedTrsBuilderFactory signed_trs.TransactionBuilderFactory
}

// CreateTransactionRepository creates a new TransactionRepository instance
func CreateTransactionRepository(
	metaDataRepository metadata.MetaDataRepository,
	sigRepository users.SignatureRepository,
	transactionRepository trs.TransactionRepository,
	signedTrsBuilderFactory signed_trs.TransactionBuilderFactory,
) signed_trs.TransactionRepository {
	out := TransactionRepository{
		metaDataRepository:      metaDataRepository,
		sigRepository:           sigRepository,
		transactionRepository:   transactionRepository,
		signedTrsBuilderFactory: signedTrsBuilderFactory,
	}
	return &out
}

// Retrieve retrieves a Transaction instance
func (rep *TransactionRepository) Retrieve(dirPath string) (signed_trs.Transaction, error) {
	//retrieve the metadata:
	met, metErr := rep.metaDataRepository.Retrieve(dirPath)
	if metErr != nil {
		return nil, metErr
	}

	//retrieve the signature:
	sigPath := filepath.Join(dirPath, "signature")
	userSig, userSigErr := rep.sigRepository.Retrieve(sigPath)
	if userSigErr != nil {
		return nil, userSigErr
	}

	//retrieve the transaction:
	trsDirPath := filepath.Join(dirPath, "transaction")
	trs, trsErr := rep.transactionRepository.Retrieve(trsDirPath)
	if trsErr != nil {
		return nil, trsErr
	}

	//build the signed transaction:
	id := met.GetID()
	ts := met.CreatedOn()
	signedTrs, signedTrsErr := rep.signedTrsBuilderFactory.Create().Create().WithID(id).WithSignature(userSig).WithTransaction(trs).CreatedOn(ts).Now()
	if signedTrsErr != nil {
		return nil, signedTrsErr
	}

	return signedTrs, nil
}

// RetrieveAll retrieves a []Transaction instances
func (rep *TransactionRepository) RetrieveAll(dirPath string) ([]signed_trs.Transaction, error) {
	files, filesErr := ioutil.ReadDir(dirPath)
	if filesErr != nil {
		return nil, filesErr
	}

	out := []signed_trs.Transaction{}
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
