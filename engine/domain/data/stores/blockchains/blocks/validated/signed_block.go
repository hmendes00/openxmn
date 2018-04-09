package validated

import (
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/users"
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
)

// SignedBlock represents a stored validated SignedBlock
type SignedBlock interface {
	GetMetaData() stored_files.File
	GetBlock() Block
	GetSignature() stored_users.Signature
}
