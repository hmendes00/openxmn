package infrastructure

import (
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	user "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
)

// User represents the concrete user
type User struct {
	Met *concrete_metadata.MetaData      `json:"metadata"`
	PK  *concrete_cryptography.PublicKey `json:"public_key"`
}

func createUser(met *concrete_metadata.MetaData, pub *concrete_cryptography.PublicKey) user.User {
	out := User{
		Met: met,
		PK:  pub,
	}

	return &out
}

// GetMetaData returns the user MetaData
func (user *User) GetMetaData() metadata.MetaData {
	return user.Met
}

// GetPublicKey returns the PublicKey
func (user *User) GetPublicKey() cryptography.PublicKey {
	return user.PK
}
