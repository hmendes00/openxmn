package infrastructure

import (
	"path/filepath"

	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	signed_trs "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/domain"
	transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
)

// AtomicTransactionService represents a concrete AtomicTransactionService implementation
type AtomicTransactionService struct {
	metaDataBuilderFactory        metadata.MetaDataBuilderFactory
	metaDataService               metadata.MetaDataService
	hashTreeService               hashtrees.HashTreeService
	userSigService                users.SignatureService
	trsService                    transactions.TransactionService
	storedAtomicTrsBuilderFactory stored_signed_transactions.AtomicTransactionBuilderFactory
}

// CreateAtomicTransactionService creates a new AtomicTransactionService instance
func CreateAtomicTransactionService(
	metaDataBuilderFactory metadata.MetaDataBuilderFactory,
	metaDataService metadata.MetaDataService,
	hashTreeService hashtrees.HashTreeService,
	userSigService users.SignatureService,
	trsService transactions.TransactionService,
	storedAtomicTrsBuilderFactory stored_signed_transactions.AtomicTransactionBuilderFactory,
) signed_trs.AtomicTransactionService {
	out := AtomicTransactionService{
		metaDataBuilderFactory:        metaDataBuilderFactory,
		metaDataService:               metaDataService,
		hashTreeService:               hashTreeService,
		userSigService:                userSigService,
		trsService:                    trsService,
		storedAtomicTrsBuilderFactory: storedAtomicTrsBuilderFactory,
	}
	return &out
}

// Save save a signed AtomicTransaction on disk
func (serv *AtomicTransactionService) Save(dirPath string, atomicTrs signed_trs.AtomicTransaction) (stored_signed_transactions.AtomicTransaction, error) {
	//build the metadata:
	id := atomicTrs.GetID()
	ts := atomicTrs.CreatedOn()
	met, metErr := serv.metaDataBuilderFactory.Create().Create().WithID(id).CreatedOn(ts).Now()
	if metErr != nil {
		return nil, metErr
	}

	//save the metadata:
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//save the hashtree:
	ht := atomicTrs.GetHashTree()
	storedHt, storedHtErr := serv.hashTreeService.Save(dirPath, ht)
	if storedHtErr != nil {
		return nil, storedHtErr
	}

	//save the signature:
	sig := atomicTrs.GetSignature()
	storedSig, storedSigErr := serv.userSigService.Save(dirPath, sig)
	if storedSigErr != nil {
		return nil, storedSigErr
	}

	//save the transactions:
	trs := atomicTrs.GetTrs()
	trsDirPath := filepath.Join(dirPath, "transactions")
	storedTrs, storedTrsErr := serv.trsService.SaveAll(trsDirPath, trs)
	if storedTrsErr != nil {
		return nil, storedTrsErr
	}

	//build the stored atomic transaction:
	storedAtomicTrs, storedAtomicTrsErr := serv.storedAtomicTrsBuilderFactory.Create().Create().WithHashTree(storedHt).WithMetaData(storedMet).WithSignature(storedSig).WithTransactions(storedTrs).Now()
	if storedAtomicTrsErr != nil {
		return nil, storedAtomicTrsErr
	}

	return storedAtomicTrs, nil
}

// SaveAll saves []AtomicTransaction on disk
func (serv *AtomicTransactionService) SaveAll(dirPath string, atomicTrs []signed_trs.AtomicTransaction) ([]stored_signed_transactions.AtomicTransaction, error) {
	out := []stored_signed_transactions.AtomicTransaction{}
	for _, oneAtomicTrs := range atomicTrs {
		oneAtomicTrsDirPath := filepath.Join(dirPath, oneAtomicTrs.GetID().String())
		oneAtomicTrs, oneAtomicTrsErr := serv.Save(oneAtomicTrsDirPath, oneAtomicTrs)
		if oneAtomicTrsErr != nil {
			return nil, oneAtomicTrsErr
		}

		out = append(out, oneAtomicTrs)
	}

	return out, nil
}
