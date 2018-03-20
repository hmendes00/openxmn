package users

import (
	"path/filepath"

	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/files"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/users"
)

// SignaturesRepository represents a concrete stored signatures repository
type SignaturesRepository struct {
	fileRepository           stored_files.FileRepository
	signatureRepository      stored_users.SignatureRepository
	signaturesBuilderFactory stored_users.SignaturesBuilderFactory
}

// CreateSignaturesRepository creates a new SignaturesRepository instance
func CreateSignaturesRepository(fileRepository stored_files.FileRepository, signatureRepository stored_users.SignatureRepository, signaturesBuilderFactory stored_users.SignaturesBuilderFactory) stored_users.SignaturesRepository {
	out := SignaturesRepository{
		fileRepository:           fileRepository,
		signatureRepository:      signatureRepository,
		signaturesBuilderFactory: signaturesBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a stored signatures
func (rep *SignaturesRepository) Retrieve(dirPath string) (stored_users.Signatures, error) {
	metFilePath := filepath.Join(dirPath, "metadata.json")
	met, metErr := rep.fileRepository.Retrieve(metFilePath)
	if metErr != nil {
		return nil, metErr
	}

	sigs, sigsErr := rep.signatureRepository.RetrieveAll(dirPath)
	if sigsErr != nil {
		return nil, sigsErr
	}

	out, outErr := rep.signaturesBuilderFactory.Create().Create().WithMetaData(met).WithSignatures(sigs).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
