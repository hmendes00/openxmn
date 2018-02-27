package infrastructure

import (
	"path/filepath"

	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	signed_trs "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/domain"
	trs "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	stored_signed_transaction "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
)

// TransactionService represents a concrete TransactionService implementation
type TransactionService struct {
	metaDataBuilderFactory        metadata.MetaDataBuilderFactory
	metaDataService               metadata.MetaDataService
	trsService                    trs.TransactionService
	storedSignedTrsBuilderFactory stored_signed_transaction.TransactionBuilderFactory
	sigService                    users.SignatureService
}

// CreateTransactionService creates a new TransactionService instance
func CreateTransactionService(
	metaDataBuilderFactory metadata.MetaDataBuilderFactory,
	metaDataService metadata.MetaDataService,
	trsService trs.TransactionService,
	storedSignedTrsBuilderFactory stored_signed_transaction.TransactionBuilderFactory,
	sigService users.SignatureService,
) signed_trs.TransactionService {
	out := TransactionService{
		metaDataBuilderFactory:        metaDataBuilderFactory,
		metaDataService:               metaDataService,
		trsService:                    trsService,
		storedSignedTrsBuilderFactory: storedSignedTrsBuilderFactory,
		sigService:                    sigService,
	}
	return &out
}

// Save save a signed Transaction on disk
func (serv *TransactionService) Save(dirPath string, signedTrs signed_trs.Transaction) (stored_signed_transaction.Transaction, error) {
	//build the metadata:
	id := signedTrs.GetID()
	createdOn := signedTrs.CreatedOn()
	met, metErr := serv.metaDataBuilderFactory.Create().Create().WithID(id).CreatedOn(createdOn).Now()
	if metErr != nil {
		return nil, metErr
	}

	//save the metadata:
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//save the signature:
	sig := signedTrs.GetSignature()
	storedSig, storedSigErr := serv.sigService.Save(dirPath, sig)
	if storedSigErr != nil {
		return nil, storedSigErr
	}

	//save the transaction:
	trs := signedTrs.GetTrs()
	trsPath := filepath.Join(dirPath, "transaction")
	storedTrsObj, storedTrsObjErr := serv.trsService.Save(trsPath, trs)
	if storedTrsObjErr != nil {
		return nil, storedTrsObjErr
	}

	//build the stored signed transaction:
	storedSignedTrs, storedSignedTrsErr := serv.storedSignedTrsBuilderFactory.Create().Create().WithMetaData(storedMet).WithSignature(storedSig).WithTransaction(storedTrsObj).Now()
	if storedSignedTrsErr != nil {
		return nil, storedSignedTrsErr
	}

	return storedSignedTrs, nil
}

// SaveAll saves signed []Transaction on disk
func (serv *TransactionService) SaveAll(dirPath string, signedTrs []signed_trs.Transaction) ([]stored_signed_transaction.Transaction, error) {
	out := []stored_signed_transaction.Transaction{}
	for _, oneSignedTrs := range signedTrs {
		signedTrsPath := filepath.Join(dirPath, oneSignedTrs.GetID().String())
		oneTrs, oneTrsErr := serv.Save(signedTrsPath, oneSignedTrs)
		if oneTrsErr != nil {
			return nil, oneTrsErr
		}

		out = append(out, oneTrs)
	}

	return out, nil
}
