package safes

import (
	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
)

// Safe represents a safe of tokens
type Safe interface {
	GetMetaData() metadata.MetaData
	GetToken() tokens.Token
	GetCipher() cryptography.Cipher
}
