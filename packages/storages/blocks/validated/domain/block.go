package domain

import (
	stored_block "github.com/XMNBlockchain/core/packages/storages/blocks/blocks/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_users "github.com/XMNBlockchain/core/packages/storages/users/domain"
)

// Block represents a stored validated block
type Block interface {
	GetMetaData() stored_files.File
	GetBlock() stored_block.SignedBlock
	GetSignatures() stored_users.Signatures
}
