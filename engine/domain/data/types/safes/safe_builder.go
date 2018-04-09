package safes

import (
	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
)

// SafeBuilder represents a safe builder
type SafeBuilder interface {
	Create() SafeBuilder
	WithMetaData(met metadata.MetaData) SafeBuilder
	WithToken(tok tokens.Token) SafeBuilder
	WithCipher(cipher cryptography.Cipher) SafeBuilder
	Now() (Safe, error)
}
