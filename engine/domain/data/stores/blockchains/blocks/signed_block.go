package blocks

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/users"
)

// SignedBlock represents a stored signed block
type SignedBlock interface {
	GetMetaData() stored_files.File
	GetSignature() stored_users.Signature
	GetBlock() Block
}
