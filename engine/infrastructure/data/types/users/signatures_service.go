package users

import (
	"path/filepath"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/users"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// SignaturesService represents a concrete SignaturesService implementation
type SignaturesService struct {
	metaDataService          metadata.MetaDataService
	sigService               users.SignatureService
	storedSigsBuilderFactory stored_users.SignaturesBuilderFactory
}

// CreateSignaturesService creates a new SignaturesService instance
func CreateSignaturesService(metaDataService metadata.MetaDataService, sigService users.SignatureService, storedSigsBuilderFactory stored_users.SignaturesBuilderFactory) users.SignaturesService {
	out := SignaturesService{
		metaDataService:          metaDataService,
		sigService:               sigService,
		storedSigsBuilderFactory: storedSigsBuilderFactory,
	}

	return &out
}

// Save saves a Signatures instance
func (serv *SignaturesService) Save(dirPath string, sig users.Signatures) (stored_users.Signatures, error) {
	//save the metadata:
	met := sig.GetMetaData()
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//save the signatures:
	sigs := sig.GetSignatures()
	sigsPath := filepath.Join(dirPath, "signatures")
	storedSigs, storedSigsErr := serv.sigService.SaveAll(sigsPath, sigs)
	if storedSigsErr != nil {
		return nil, storedSigsErr
	}

	//build the stored signatures:
	storedUserSigs, storedUserSigsErr := serv.storedSigsBuilderFactory.Create().Create().WithMetaData(storedMet).WithSignatures(storedSigs).Now()
	if storedUserSigsErr != nil {
		return nil, storedUserSigsErr
	}

	return storedUserSigs, nil
}
