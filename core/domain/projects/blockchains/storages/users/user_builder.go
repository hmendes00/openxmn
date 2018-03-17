package users

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
)

// UserBuilder represents a stored user builder
type UserBuilder interface {
	Create() UserBuilder
	WithMetaData(met stored_files.File) UserBuilder
	WithPublicKey(pk stored_files.File) UserBuilder
	Now() (User, error)
}
