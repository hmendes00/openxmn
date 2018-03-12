package users

import (
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/users"
)

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
