package infrastructure

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_users "github.com/XMNBlockchain/core/packages/storages/users/domain"
)

type user struct {
	met    stored_files.File
	pubKey stored_files.File
}

func createUser(met stored_files.File, pubKey stored_files.File) stored_users.User {
	out := user{
		met:    met,
		pubKey: pubKey,
	}

	return &out
}

// GetMetaData returns the MetaData
func (usr *user) GetMetaData() stored_files.File {
	return usr.met
}

// GetPublicKey returns the Public Key
func (usr *user) GetPublicKey() stored_files.File {
	return usr.pubKey
}
