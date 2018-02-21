package infrastructure

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
	users "github.com/XMNBlockchain/core/packages/lives/users/domain"
)

// SignatureRepository represents a concrete SignatureRepository implementation
type SignatureRepository struct {
	fileRepository files.FileRepository
}

// CreateSignatureRepository creates a new SignatureRepository instance
func CreateSignatureRepository(fileRepository files.FileRepository) users.SignatureRepository {
	out := SignatureRepository{
		fileRepository: fileRepository,
	}
	return &out
}

// Retrieve retrieves a Signature instance
func (rep *SignatureRepository) Retrieve(dirPath string) (users.Signature, error) {
	fil, filErr := rep.fileRepository.Retrieve(dirPath, "user_signature.json")
	if filErr != nil {
		return nil, filErr
	}

	//convert the data to a signature instance:
	newSig := new(Signature)
	jsonErr := json.Unmarshal(fil.GetData(), newSig)
	if jsonErr != nil {
		return nil, jsonErr
	}

	//return the signature:
	return newSig, nil
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
