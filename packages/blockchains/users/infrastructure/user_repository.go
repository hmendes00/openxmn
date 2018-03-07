package infrastructure

import (
	files "github.com/XMNBlockchain/core/packages/blockchains/files/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
)

// UserRepository represents a concrete UserRepository implementation
type UserRepository struct {
	metaDataRepository   metadata.MetaDataRepository
	fileRepository       files.FileRepository
	pubKeyBuilderFactory cryptography.PublicKeyBuilderFactory
	usrBuilderFactory    users.UserBuilderFactory
}

// CreateUserRepository creates a new UserRepository instance
func CreateUserRepository(metaDataRepository metadata.MetaDataRepository, fileRepository files.FileRepository, pubKeyBuilderFactory cryptography.PublicKeyBuilderFactory, usrBuilderFactory users.UserBuilderFactory) users.UserRepository {
	out := UserRepository{
		metaDataRepository:   metaDataRepository,
		fileRepository:       fileRepository,
		pubKeyBuilderFactory: pubKeyBuilderFactory,
		usrBuilderFactory:    usrBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a User instance
func (rep *UserRepository) Retrieve(dirPath string) (users.User, error) {
	//retrieve the metadata:
	met, metErr := rep.metaDataRepository.Retrieve(dirPath)
	if metErr != nil {
		return nil, metErr
	}

	//retrieve the public key file:
	pubKeyFile, pubKeyFileErr := rep.fileRepository.Retrieve(dirPath, "key.pub")
	if pubKeyFileErr != nil {
		return nil, pubKeyFileErr
	}

	//build the public key:
	pubKeyAsBytes := pubKeyFile.GetData()
	pubKey, pubKeyErr := rep.pubKeyBuilderFactory.Create().Create().WithEncodedString(string(pubKeyAsBytes)).Now()
	if pubKeyErr != nil {
		return nil, pubKeyErr
	}

	//build the user:
	usr, usrErr := rep.usrBuilderFactory.Create().Create().WithMetaData(met).WithPublicKey(pubKey).Now()
	if usrErr != nil {
		return nil, usrErr
	}

	return usr, nil
}
