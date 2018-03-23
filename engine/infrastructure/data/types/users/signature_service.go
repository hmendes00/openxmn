package users

import (
	"encoding/json"
	"path/filepath"

	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/files"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/users"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// SignatureService represents a concrete SignatureService implementation
type SignatureService struct {
	metaDataService         metadata.MetaDataService
	usrService              users.UserService
	fileService             files.FileService
	fileBuilderFactory      files.FileBuilderFactory
	storedSigBuilderFactory stored_users.SignatureBuilderFactory
}

// CreateSignatureService creates a new SignatureService instance
func CreateSignatureService(metaDataService metadata.MetaDataService, usrService users.UserService, fileService files.FileService, fileBuilderFactory files.FileBuilderFactory, storedSigBuilderFactory stored_users.SignatureBuilderFactory) users.SignatureService {
	out := SignatureService{
		metaDataService:         metaDataService,
		usrService:              usrService,
		fileService:             fileService,
		fileBuilderFactory:      fileBuilderFactory,
		storedSigBuilderFactory: storedSigBuilderFactory,
	}
	return &out
}

// Save saves a Signature
func (serv *SignatureService) Save(dirPath string, sig users.Signature) (stored_users.Signature, error) {
	//save the metadata:
	met := sig.GetMetaData()
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//convert the signature to json:
	sigSig := sig.GetSignature()
	sigAsJS, sigAsJSErr := json.Marshal(sigSig)
	if sigAsJSErr != nil {
		return nil, sigAsJSErr
	}

	//build the user signature file:
	sigFile, sigFileErr := serv.fileBuilderFactory.Create().Create().WithData(sigAsJS).WithFileName("signature").WithExtension("json").Now()
	if sigFileErr != nil {
		return nil, sigFileErr
	}

	//save the user signature:
	storedSigFile, storedSigFileErr := serv.fileService.Save(dirPath, sigFile)
	if storedSigFileErr != nil {
		return nil, storedSigFileErr
	}

	//save the user:
	usr := sig.GetUser()
	usrFilePath := filepath.Join(dirPath, "user")
	storedUser, storedUserErr := serv.usrService.Save(usrFilePath, usr)
	if storedUserErr != nil {
		return nil, storedUserErr
	}

	//build the stored signature:
	storedSig, storedSigErr := serv.storedSigBuilderFactory.Create().Create().WithMetaData(storedMet).WithSignature(storedSigFile).WithUser(storedUser).Now()
	if storedSigErr != nil {
		return nil, storedSigErr
	}

	return storedSig, nil
}

// SaveAll saves []Signature instances
func (serv *SignatureService) SaveAll(dirPath string, sigs []users.Signature) ([]stored_users.Signature, error) {
	out := []stored_users.Signature{}
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
