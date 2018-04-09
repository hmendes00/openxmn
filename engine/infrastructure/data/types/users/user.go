package users

import (
	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	user "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	concrete_cryptography "github.com/XMNBlockchain/openxmn/engine/infrastructure/cryptography"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
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
