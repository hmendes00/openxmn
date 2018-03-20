package users

import (
	"io/ioutil"
	"path/filepath"

	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/users"
)

// SignatureRepository represents a concrete stored signature repository
type SignatureRepository struct {
	fileRepository    stored_files.FileRepository
	userRepository    stored_users.UserRepository
	sigBuilderFactory stored_users.SignatureBuilderFactory
}

// CreateSignatureRepository creates a new SignatureRepository instance
func CreateSignatureRepository(fileRepository stored_files.FileRepository, userRepository stored_users.UserRepository, sigBuilderFactory stored_users.SignatureBuilderFactory) stored_users.SignatureRepository {
	out := SignatureRepository{
		fileRepository:    fileRepository,
		userRepository:    userRepository,
		sigBuilderFactory: sigBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a stored signature
func (rep *SignatureRepository) Retrieve(dirPath string) (stored_users.Signature, error) {
	metFilePath := filepath.Join(dirPath, "metadata.json")
	met, metErr := rep.fileRepository.Retrieve(metFilePath)
	if metErr != nil {
		return nil, metErr
	}

	sigFilePath := filepath.Join(dirPath, "signature.json")
	sig, sigErr := rep.fileRepository.Retrieve(sigFilePath)
	if sigErr != nil {
		return nil, sigErr
	}

	usrFilePath := filepath.Join(dirPath, "user")
	usr, usrErr := rep.userRepository.Retrieve(usrFilePath)
	if usrErr != nil {
		return nil, usrErr
	}

	out, outErr := rep.sigBuilderFactory.Create().Create().WithMetaData(met).WithSignature(sig).WithUser(usr).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}

// RetrieveAll retrieves stored []Signature
func (rep *SignatureRepository) RetrieveAll(dirPath string) ([]stored_users.Signature, error) {
	fileInfs, fileInfsErr := ioutil.ReadDir(dirPath)
	if fileInfsErr != nil {
		return nil, fileInfsErr
	}

	out := []stored_users.Signature{}
	for _, oneFileInf := range fileInfs {
		if oneFileInf.IsDir() {
			continue
		}

		sigPath := filepath.Join(dirPath, oneFileInf.Name())
		sig, sigErr := rep.Retrieve(sigPath)
		if sigErr != nil {
			continue
		}

		out = append(out, sig)
	}

	return out, nil
}
