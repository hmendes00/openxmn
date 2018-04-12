package safes

import (
	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
)

// Safe represents a safe of tokens
type Safe interface {
	GetMetaData() metadata.MetaData
	GetCipher() cryptography.Cipher
}
