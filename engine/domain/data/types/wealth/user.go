package wealth

import (
	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	types "github.com/XMNBlockchain/openxmn/engine/domain/data/types"
)

// User represents a user
type User interface {
	GetMetaData() types.MetaData
	GetPublicKey() cryptography.PublicKey
}
