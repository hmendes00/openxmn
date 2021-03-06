package users

import (
	"time"

	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	user "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	concrete_cryptography "github.com/XMNBlockchain/openxmn/engine/infrastructure/cryptography"
	concrete_stored_user "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/users"
	concrete_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/files"
	concrete_hashtrees "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/hashtrees"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	uuid "github.com/satori/go.uuid"
)

// CreateUserForTests creates a User for tests
func CreateUserForTests() *User {
	pk := concrete_cryptography.CreatePublicKeyForTests()
	return CreateUserUsingProvidedPublicKeyForTests(pk)
}

// CreateUserUsingProvidedPublicKeyForTests creates a User with a provided PublicKey for tests
func CreateUserUsingProvidedPublicKeyForTests(pk cryptography.PublicKey) *User {
	//variables:
	id := uuid.NewV4()
	crOn := time.Now().UTC()
	met, _ := concrete_metadata.CreateBuilderFactory().Create().Create().WithID(&id).CreatedOn(crOn).Now()

	user := createUser(met.(*concrete_metadata.MetaData), pk.(*concrete_cryptography.PublicKey))
	return user.(*User)
}

// CreateSignatureForTests creates a Signature for tests
func CreateSignatureForTests() *Signature {
	//variables:
	id := uuid.NewV4()
	sig := concrete_cryptography.CreateSignatureForTests()
	usr := CreateUserUsingProvidedPublicKeyForTests(sig.GetPublicKey())
	crOn := time.Now().UTC()

	met, _ := concrete_metadata.CreateBuilderFactory().Create().Create().WithID(&id).CreatedOn(crOn).Now()

	userSig, _ := createSignature(met.(*concrete_metadata.MetaData), sig, usr)
	return userSig.(*Signature)
}

// CreateSignaturesForTests creates a Signature for tests
func CreateSignaturesForTests() *Signatures {
	//variables:
	id := uuid.NewV4()
	sigs := []*Signature{
		CreateSignatureForTests(),
		CreateSignatureForTests(),
		CreateSignatureForTests(),
	}

	crOn := time.Now().UTC()
	met, _ := concrete_metadata.CreateBuilderFactory().Create().Create().WithID(&id).CreatedOn(crOn).Now()

	userSigs := createSignatures(met.(*concrete_metadata.MetaData), sigs)
	return userSigs.(*Signatures)
}

// CreateUserBuilderFactoryForTests creates a UserBuilderFactory for tests
func CreateUserBuilderFactoryForTests() user.UserBuilderFactory {
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactoryForTests()
	metaDataBuilderFactory := concrete_metadata.CreateBuilderFactoryForTests()
	out := CreateUserBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	return out
}

// CreateUserRepositoryForTests creates a UserRepository for tests
func CreateUserRepositoryForTests() user.UserRepository {
	metaDataRepository := concrete_metadata.CreateRepositoryForTests()
	fileRepository := concrete_files.CreateFileRepositoryForTests()
	pubKeyBuilderFactory := concrete_cryptography.CreatePublicKeyBuilderFactoryForTests()
	usrBuilderFactory := CreateUserBuilderFactoryForTests()
	out := CreateUserRepository(metaDataRepository, fileRepository, pubKeyBuilderFactory, usrBuilderFactory)
	return out
}

// CreateUserServiceForTests creates a UserService for tests
func CreateUserServiceForTests() user.UserService {
	metaDataService := concrete_metadata.CreateServiceForTests()
	fileService := concrete_files.CreateFileServiceForTests()
	fileBuilderFactory := concrete_files.CreateFileBuilderFactoryForTests()
	storedUserBuilderFactory := concrete_stored_user.CreateUserBuilderFactoryForTests()
	out := CreateUserService(metaDataService, fileService, fileBuilderFactory, storedUserBuilderFactory)
	return out
}

// CreateSignatureBuilderFactoryForTests creates a new SignatureBuilderFactory for tests
func CreateSignatureBuilderFactoryForTests() user.SignatureBuilderFactory {
	sigBuilderFactory := concrete_cryptography.CreateSignatureBuilderFactoryForTests()
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactoryForTests()
	metaDataBuilderFactory := concrete_metadata.CreateBuilderFactoryForTests()
	out := CreateSignatureBuilderFactory(sigBuilderFactory, htBuilderFactory, metaDataBuilderFactory)
	return out
}

// CreateSignatureRepositoryForTests creates a new SignatureRepository for tests
func CreateSignatureRepositoryForTests() user.SignatureRepository {
	metaDataRepository := concrete_metadata.CreateRepositoryForTests()
	usrRepository := CreateUserRepositoryForTests()
	fileRepository := concrete_files.CreateFileRepositoryForTests()
	userSigBuilderFactory := CreateSignatureBuilderFactoryForTests()
	out := CreateSignatureRepository(metaDataRepository, usrRepository, fileRepository, userSigBuilderFactory)
	return out
}

// CreateSignatureServiceForTests creates a new SignatureService for tests
func CreateSignatureServiceForTests() user.SignatureService {
	metaDataService := concrete_metadata.CreateServiceForTests()
	usrService := CreateUserServiceForTests()
	fileService := concrete_files.CreateFileServiceForTests()
	fileBuilderFactory := concrete_files.CreateFileBuilderFactoryForTests()
	storedSigBuilderFactory := concrete_stored_user.CreateSignatureBuilderFactoryForTests()
	out := CreateSignatureService(metaDataService, usrService, fileService, fileBuilderFactory, storedSigBuilderFactory)
	return out
}

// CreateSignaturesBuilderFactoryForTests creates a new SignaturesBuilderFactory for tests
func CreateSignaturesBuilderFactoryForTests() user.SignaturesBuilderFactory {
	metaDataBuilderFactory := concrete_metadata.CreateBuilderFactoryForTests()
	out := CreateSignaturesBuilderFactory(metaDataBuilderFactory)
	return out
}

// CreateSignaturesRepositoryForTests creates a new SignaturesRepository for tests
func CreateSignaturesRepositoryForTests() user.SignaturesRepository {
	metaDataRepository := concrete_metadata.CreateRepositoryForTests()
	sigsRepository := CreateSignatureRepositoryForTests()
	sigsBuilderFactory := CreateSignaturesBuilderFactoryForTests()
	out := CreateSignaturesRepository(metaDataRepository, sigsRepository, sigsBuilderFactory)
	return out
}

// CreateSignaturesServiceForTests creates a new SignaturesService for tests
func CreateSignaturesServiceForTests() user.SignaturesService {
	metaDataService := concrete_metadata.CreateServiceForTests()
	sigService := CreateSignatureServiceForTests()
	storedSigsBuilderFactory := concrete_stored_user.CreateSignaturesBuilderFactoryForTests()
	out := CreateSignaturesService(metaDataService, sigService, storedSigsBuilderFactory)
	return out
}
