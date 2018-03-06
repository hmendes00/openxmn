package infrastructure

import (
	"path/filepath"

	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	signed_trs "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/domain"
	transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
)

// AtomicTransactionService represents a concrete AtomicTransactionService implementation
type AtomicTransactionService struct {
	metaDataService               metadata.MetaDataService
	userSigService                users.SignatureService
	trsService                    transactions.TransactionsService
	storedAtomicTrsBuilderFactory stored_signed_transactions.AtomicTransactionBuilderFactory
}

// CreateAtomicTransactionService creates a new AtomicTransactionService instance
func CreateAtomicTransactionService(
	metaDataService metadata.MetaDataService,
	userSigService users.SignatureService,
	trsService transactions.TransactionsService,
	storedAtomicTrsBuilderFactory stored_signed_transactions.AtomicTransactionBuilderFactory,
) signed_trs.AtomicTransactionService {
	out := AtomicTransactionService{
		metaDataService:               metaDataService,
		userSigService:                userSigService,
		trsService:                    trsService,
		storedAtomicTrsBuilderFactory: storedAtomicTrsBuilderFactory,
	}
	return &out
}

// Save save a signed AtomicTransaction on disk
func (serv *AtomicTransactionService) Save(dirPath string, atomicTrs signed_trs.AtomicTransaction) (stored_signed_transactions.AtomicTransaction, error) {
	//save the metadata:
	met := atomicTrs.GetMetaData()
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//save the signature:
	sig := atomicTrs.GetSignature()
	storedSig, storedSigErr := serv.userSigService.Save(dirPath, sig)
	if storedSigErr != nil {
		return nil, storedSigErr
	}

	//save the transactions:
	trs := atomicTrs.GetTransactions()
	trsDirPath := filepath.Join(dirPath, "transactions")
	storedTrs, storedTrsErr := serv.trsService.Save(trsDirPath, trs)
	if storedTrsErr != nil {
		return nil, storedTrsErr
	}

	//build the stored atomic transaction:
	storedAtomicTrs, storedAtomicTrsErr := serv.storedAtomicTrsBuilderFactory.Create().Create().WithMetaData(storedMet).WithSignature(storedSig).WithTransactions(storedTrs).Now()
	if storedAtomicTrsErr != nil {
		return nil, storedAtomicTrsErr
	}

	return storedAtomicTrs, nil
}

// SaveAll saves []AtomicTransaction on disk
func (serv *AtomicTransactionService) SaveAll(dirPath string, atomicTrs []signed_trs.AtomicTransaction) ([]stored_signed_transactions.AtomicTransaction, error) {
	out := []stored_signed_transactions.AtomicTransaction{}
	for _, oneAtomicTrs := range atomicTrs {
		oneAtomicTrsDirPath := filepath.Join(dirPath, oneAtomicTrs.GetMetaData().GetID().String())
		oneAtomicTrs, oneAtomicTrsErr := serv.Save(oneAtomicTrsDirPath, oneAtomicTrs)
		if oneAtomicTrsErr != nil {
			return nil, oneAtomicTrsErr
		}

		out = append(out, oneAtomicTrs)
	}

	return out, nil
}
