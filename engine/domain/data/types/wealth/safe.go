package wealth

import (
	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	types "github.com/XMNBlockchain/openxmn/engine/domain/data/types"
)

// Safe represents a safe containing a specific amount of token
type Safe interface {
	GetMetaData() types.MetaData
	GetToken() Token
	GetAmount() cryptography.Cipher
}
