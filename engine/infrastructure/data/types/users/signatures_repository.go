package users

import (
	"path/filepath"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// SignaturesRepository represents a concrete SignaturesRepository implementation
type SignaturesRepository struct {
	metaDataRepository metadata.Repository
	sigsRepository     users.SignatureRepository
	sigsBuilderFactory users.SignaturesBuilderFactory
}

// CreateSignaturesRepository creates a new SignaturesRepository instance
func CreateSignaturesRepository(metaDataRepository metadata.Repository, sigsRepository users.SignatureRepository, sigsBuilderFactory users.SignaturesBuilderFactory) users.SignaturesRepository {
	out := SignaturesRepository{
		metaDataRepository: metaDataRepository,
		sigsRepository:     sigsRepository,
		sigsBuilderFactory: sigsBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a Signatures instance
func (rep *SignaturesRepository) Retrieve(dirPath string) (users.Signatures, error) {
	//retrieve the metadata:
	met, metErr := rep.metaDataRepository.Retrieve(dirPath)
	if metErr != nil {
		return nil, metErr
	}

	//retrieve the signatures:
	sigsPath := filepath.Join(dirPath, "signatures")
	sigs, sigsErr := rep.sigsRepository.RetrieveAll(sigsPath)
	if sigsErr != nil {
		return nil, sigsErr
	}

	//build the user signatures:
	userSigs, userSigsErr := rep.sigsBuilderFactory.Create().Create().WithMetaData(met).WithSignatures(sigs).Now()
	if userSigsErr != nil {
		return nil, userSigsErr
	}

	return userSigs, nil
}
