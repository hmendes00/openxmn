package users

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/users"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
)

// User represents a concrete stored user implementation
type User struct {
	Met    *concrete_stored_files.File `json:"metadata"`
	PubKey *concrete_stored_files.File `json:"public_key"`
}

func createUser(met *concrete_stored_files.File, pubKey *concrete_stored_files.File) stored_users.User {
	out := User{
		Met:    met,
		PubKey: pubKey,
	}

	return &out
}

// GetMetaData returns the MetaData
func (usr *User) GetMetaData() stored_files.File {
	return usr.Met
}

// GetPublicKey returns the Public Key
func (usr *User) GetPublicKey() stored_files.File {
	return usr.PubKey
}
