package transactions

import (
	"io/ioutil"
	"path/filepath"

	chunks "github.com/XMNBlockchain/openxmn/engine/domain/data/types/chunks"
	met "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	trs "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
)

// TransactionRepository represents a concrete TransactionRepository implementation
type TransactionRepository struct {
	chksRepository            chunks.Repository
	metRepository             met.Repository
	transactionBuilderFactory trs.TransactionBuilderFactory
}

// CreateTransactionRepository creates a new TransactionRepository instance
func CreateTransactionRepository(
	chksRepository chunks.Repository,
	metRepository met.Repository,
	transactionBuilderFactory trs.TransactionBuilderFactory,
) trs.TransactionRepository {
	out := TransactionRepository{
		chksRepository:            chksRepository,
		metRepository:             metRepository,
		transactionBuilderFactory: transactionBuilderFactory,
	}
	return &out
}

// Retrieve retrieves a Transaction instance
func (rep *TransactionRepository) Retrieve(dirPath string) (trs.Transaction, error) {

	//retrieve the metadata:
	met, metErr := rep.metRepository.Retrieve(dirPath)
	if metErr != nil {
		return nil, metErr
	}

	//retrieve the chunks:
	chkDirPath := filepath.Join(dirPath, "json")
	chks, chksErr := rep.chksRepository.Retrieve(chkDirPath)
	if chksErr != nil {
		return nil, chksErr
	}

	//get the json data:
	jsData, jsDataErr := chks.GetData()
	if jsDataErr != nil {
		return nil, jsDataErr
	}

	//build the transaction:
	trs, trsErr := rep.transactionBuilderFactory.Create().Create().WithMetaData(met).WithJSON(jsData).Now()
	if trsErr != nil {
		return nil, trsErr
	}

	return trs, nil
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
