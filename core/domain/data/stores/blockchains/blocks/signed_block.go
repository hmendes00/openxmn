package blocks

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/files"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/users"
)

// SignedBlock represents a stored signed block
type SignedBlock interface {
	GetMetaData() stored_files.File
	GetSignature() stored_users.Signature
	GetBlock() Block
}
