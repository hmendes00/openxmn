package users

import (
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/users"
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/files"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// UserService represents a concrete UserService implementation
type UserService struct {
	metaDataService          metadata.Service
	fileService              files.FileService
	fileBuilderFactory       files.FileBuilderFactory
	storedUserBuilderFactory stored_users.UserBuilderFactory
}

// CreateUserService creates a new UserService instance
func CreateUserService(metaDataService metadata.Service, fileService files.FileService, fileBuilderFactory files.FileBuilderFactory, storedUserBuilderFactory stored_users.UserBuilderFactory) users.UserService {
	out := UserService{
		metaDataService:          metaDataService,
		fileService:              fileService,
		fileBuilderFactory:       fileBuilderFactory,
		storedUserBuilderFactory: storedUserBuilderFactory,
	}

	return &out
}

// Save saves a User instance
func (serv *UserService) Save(dirPath string, usr users.User) (stored_users.User, error) {
	//save the metadata:
	met := usr.GetMetaData()
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//get the public key as string:
	pubKeyAsString, pubKeyAsStringErr := usr.GetPublicKey().String()
	if pubKeyAsStringErr != nil {
		return nil, pubKeyAsStringErr
	}

	//build the public key file:
	pubKeyFile, pubKeyFileErr := serv.fileBuilderFactory.Create().Create().WithData([]byte(pubKeyAsString)).WithFileName("key").WithExtension("pub").Now()
	if pubKeyFileErr != nil {
		return nil, pubKeyFileErr
	}

	//save the public key file:
	storedPubKey, storedPubKeyErr := serv.fileService.Save(dirPath, pubKeyFile)
	if storedPubKeyErr != nil {
		return nil, storedPubKeyErr
	}

	//build the stored user:
	storedUsr, storedUsrErr := serv.storedUserBuilderFactory.Create().Create().WithMetaData(storedMet).WithPublicKey(storedPubKey).Now()
	if storedUsrErr != nil {
		return nil, storedUsrErr
	}

	return storedUsr, nil
}
