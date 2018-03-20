package users

import (
	cryptography "github.com/XMNBlockchain/exmachina-network/core/domain/cryptography"
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/metadata"
	user "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/users"
	concrete_cryptography "github.com/XMNBlockchain/exmachina-network/core/infrastructure/cryptography/rsa"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/metadata"
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
