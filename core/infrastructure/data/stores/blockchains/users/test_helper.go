package users

import (
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/users"
	concrete_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/stores/blockchains/files"
)

// CreateUserForTests creates a User for tests
func CreateUserForTests() *User {
	met := concrete_files.CreateFileForTests()
	pubKey := concrete_files.CreateFileForTests()
	out := createUser(met, pubKey)
	return out.(*User)
}

// CreateSignatureForTests creates a Signature for tests
func CreateSignatureForTests() *Signature {
	met := concrete_files.CreateFileForTests()
	sig := concrete_files.CreateFileForTests()
	usr := CreateUserForTests()
	out := createSignature(met, sig, usr)
	return out.(*Signature)
}

// CreateSignaturesForTests creates a Signatures for tests
func CreateSignaturesForTests() *Signatures {
	met := concrete_files.CreateFileForTests()
	sigs := []*Signature{
		CreateSignatureForTests(),
		CreateSignatureForTests(),
		CreateSignatureForTests(),
	}

	out := createSignatures(met, sigs)
	return out.(*Signatures)
}

// CreateUserBuilderFactoryForTests creates a new UserBuilderFactory for tests
func CreateUserBuilderFactoryForTests() stored_users.UserBuilderFactory {
	out := CreateUserBuilderFactory()
	return out
}

// CreateSignatureBuilderFactoryForTests creates a new SignatureBuilderFactory for tests
func CreateSignatureBuilderFactoryForTests() stored_users.SignatureBuilderFactory {
	out := CreateSignatureBuilderFactory()
	return out
}

// CreateSignaturesBuilderFactoryForTests creates a new SignaturesBuilderFactory for tests
func CreateSignaturesBuilderFactoryForTests() stored_users.SignaturesBuilderFactory {
	out := CreateSignaturesBuilderFactory()
	return out
}

// CreateUserRepositoryForTests creates a UserRepository for tests
func CreateUserRepositoryForTests() stored_users.UserRepository {
	fileRepository := concrete_files.CreateFileRepositoryForTests()
	userBuilderFactory := CreateUserBuilderFactoryForTests()
	out := CreateUserRepository(fileRepository, userBuilderFactory)
	return out
}

// CreateSignatureRepositoryForTests creates a Signature for tests
func CreateSignatureRepositoryForTests() stored_users.SignatureRepository {
	fileRepository := concrete_files.CreateFileRepositoryForTests()
	userRepository := CreateUserRepositoryForTests()
	sigBuilderFactory := CreateSignatureBuilderFactoryForTests()
	out := CreateSignatureRepository(fileRepository, userRepository, sigBuilderFactory)
	return out
}

// CreateSignaturesRepositoryForTests creates a Signatures for tests
func CreateSignaturesRepositoryForTests() stored_users.SignaturesRepository {
	fileRepository := concrete_files.CreateFileRepositoryForTests()
	signatureRepository := CreateSignatureRepositoryForTests()
	signaturesBuilderFactory := CreateSignaturesBuilderFactoryForTests()
	out := CreateSignaturesRepository(fileRepository, signatureRepository, signaturesBuilderFactory)
	return out
}
