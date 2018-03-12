package domain

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/users"
)

// SignedBlock represents a stored signed block
type SignedBlock interface {
	GetMetaData() stored_files.File
	GetSignature() stored_users.Signature
	GetBlock() Block
}
