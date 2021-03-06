package signed

import (
	"path/filepath"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	stored_signed_transaction "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions/signed"
	trs "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	signed_trs "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions/signed"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// TransactionService represents a concrete TransactionService implementation
type TransactionService struct {
	metaDataService               metadata.Service
	trsService                    trs.TransactionService
	storedSignedTrsBuilderFactory stored_signed_transaction.TransactionBuilderFactory
	sigService                    users.SignatureService
}

// CreateTransactionService creates a new TransactionService instance
func CreateTransactionService(
	metaDataService metadata.Service,
	trsService trs.TransactionService,
	storedSignedTrsBuilderFactory stored_signed_transaction.TransactionBuilderFactory,
	sigService users.SignatureService,
) signed_trs.TransactionService {
	out := TransactionService{
		metaDataService:               metaDataService,
		trsService:                    trsService,
		storedSignedTrsBuilderFactory: storedSignedTrsBuilderFactory,
		sigService:                    sigService,
	}
	return &out
}

// Save save a signed Transaction on disk
func (serv *TransactionService) Save(dirPath string, signedTrs signed_trs.Transaction) (stored_signed_transaction.Transaction, error) {
	//save the metadata:
	met := signedTrs.GetMetaData()
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//save the signature:
	sig := signedTrs.GetSignature()
	sigPath := filepath.Join(dirPath, "signature")
	storedSig, storedSigErr := serv.sigService.Save(sigPath, sig)
	if storedSigErr != nil {
		return nil, storedSigErr
	}

	//save the transaction:
	trs := signedTrs.GetTransaction()
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
		signedTrsPath := filepath.Join(dirPath, oneSignedTrs.GetMetaData().GetID().String())
		oneTrs, oneTrsErr := serv.Save(signedTrsPath, oneSignedTrs)
		if oneTrsErr != nil {
			return nil, oneTrsErr
		}

		out = append(out, oneTrs)
	}

	return out, nil
}
