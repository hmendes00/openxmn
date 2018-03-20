package users

import (
	files "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/files"
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/metadata"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/users"
	users "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/users"
)

// UserService represents a concrete UserService implementation
type UserService struct {
	metaDataService          metadata.MetaDataService
	fileService              files.FileService
	fileBuilderFactory       files.FileBuilderFactory
	storedUserBuilderFactory stored_users.UserBuilderFactory
}

// CreateUserService creates a new UserService instance
func CreateUserService(metaDataService metadata.MetaDataService, fileService files.FileService, fileBuilderFactory files.FileBuilderFactory, storedUserBuilderFactory stored_users.UserBuilderFactory) users.UserService {
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
