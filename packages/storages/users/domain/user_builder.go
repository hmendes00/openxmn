package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// UserBuilder represents a stored user builder
type UserBuilder interface {
	Create() UserBuilder
	WithMetaData(met stored_files.File) UserBuilder
	WithPublicKey(pk stored_files.File) UserBuilder
	Now() (User, error)
}
