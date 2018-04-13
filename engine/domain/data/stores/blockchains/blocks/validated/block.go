package validated

import (
	stored_block "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/blocks"
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/users"
)

// Block represents a stored validated block
type Block interface {
	GetMetaData() stored_files.File
	GetBlock() stored_block.SignedBlock
	GetSignatures() stored_users.Signatures
}
