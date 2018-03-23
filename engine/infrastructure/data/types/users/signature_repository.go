package users

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/files"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	concrete_cryptography "github.com/XMNBlockchain/openxmn/engine/infrastructure/cryptography"
)

// SignatureRepository represents a concrete SignatureRepository implementation
type SignatureRepository struct {
	metaDataRepository    metadata.MetaDataRepository
	usrRepository         users.UserRepository
	fileRepository        files.FileRepository
	userSigBuilderFactory users.SignatureBuilderFactory
}

// CreateSignatureRepository creates a new SignatureRepository instance
func CreateSignatureRepository(metaDataRepository metadata.MetaDataRepository, usrRepository users.UserRepository, fileRepository files.FileRepository, userSigBuilderFactory users.SignatureBuilderFactory) users.SignatureRepository {
	out := SignatureRepository{
		metaDataRepository:    metaDataRepository,
		usrRepository:         usrRepository,
		fileRepository:        fileRepository,
		userSigBuilderFactory: userSigBuilderFactory,
	}
	return &out
}

// Retrieve retrieves a Signature instance
func (rep *SignatureRepository) Retrieve(dirPath string) (users.Signature, error) {
	//retrieve the metadata:
	met, metErr := rep.metaDataRepository.Retrieve(dirPath)
	if metErr != nil {
		return nil, metErr
	}

	//retrieve the user:
	usrFilePath := filepath.Join(dirPath, "user")
	usr, usrErr := rep.usrRepository.Retrieve(usrFilePath)
	if usrErr != nil {
		return nil, usrErr
	}

	//retrieve the signature file:
	sigFile, sigFileErr := rep.fileRepository.Retrieve(dirPath, "signature.json")
	if sigFileErr != nil {
		return nil, sigFileErr
	}

	//build the signature:
	sigAsJS := sigFile.GetData()
	sig := new(concrete_cryptography.Signature)
	jsErr := json.Unmarshal(sigAsJS, sig)
	if jsErr != nil {
		return nil, jsErr
	}

	//build the user signature:
	userSig, userSigErr := rep.userSigBuilderFactory.Create().Create().WithMetaData(met).WithUser(usr).WithSignature(sig).Now()
	if userSigErr != nil {
		return nil, userSigErr
	}

	return userSig, nil
}

// RetrieveAll retrieves []Signature instances
func (rep *SignatureRepository) RetrieveAll(dirPath string) ([]users.Signature, error) {
	files, filesErr := ioutil.ReadDir(dirPath)
	if filesErr != nil {
		return nil, filesErr
	}

	sigs := []users.Signature{}
	for _, oneFile := range files {
		if !oneFile.IsDir() {
			continue
		}

		sigDirPath := filepath.Join(dirPath, oneFile.Name())
		oneSig, oneSigErr := rep.Retrieve(sigDirPath)
		if oneSigErr != nil {
			return nil, oneSigErr
		}

		sigs = append(sigs, oneSig)
	}

	return sigs, nil
}
