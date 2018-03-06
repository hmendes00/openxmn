package infrastructure

import (
	"encoding/json"
	"path/filepath"

	files "github.com/XMNBlockchain/core/packages/blockchains/files/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	stored_file "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// SignatureService represents a concrete SignatureService implementation
type SignatureService struct {
	fileService        files.FileService
	fileBuilderFactory files.FileBuilderFactory
}

// CreateSignatureService creates a new SignatureService instance
func CreateSignatureService(fileService files.FileService, fileBuilderFactory files.FileBuilderFactory) users.SignatureService {
	out := SignatureService{
		fileService:        fileService,
		fileBuilderFactory: fileBuilderFactory,
	}
	return &out
}

// Save saves a Signature
func (serv *SignatureService) Save(dirPath string, sig users.Signature) (stored_file.File, error) {
	//convert the signature to json:
	js, jsErr := json.Marshal(sig)
	if jsErr != nil {
		return nil, jsErr
	}

	//build the user signature file:
	htSig, htSigErr := serv.fileBuilderFactory.Create().Create().WithData(js).WithFileName("user_signature").WithExtension("json").Now()
	if htSigErr != nil {
		return nil, htSigErr
	}

	//save the user signature:
	storedSig, storedSigErr := serv.fileService.Save(dirPath, htSig)
	if storedSigErr != nil {
		return nil, storedSigErr
	}

	return storedSig, nil
}

// SaveAll saves []Signature instances
func (serv *SignatureService) SaveAll(dirPath string, sigs []users.Signature) ([]stored_file.File, error) {
	out := []stored_file.File{}
	for _, oneSig := range sigs {
		oneSigPath := filepath.Join(dirPath, oneSig.GetMetaData().GetID().String())
		fil, filErr := serv.Save(oneSigPath, oneSig)
		if filErr != nil {
			return nil, filErr
		}

		out = append(out, fil)
	}

	return out, nil
}
