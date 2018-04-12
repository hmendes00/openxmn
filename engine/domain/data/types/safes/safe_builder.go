package safes

import (
	"time"

	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	uuid "github.com/satori/go.uuid"
)

// SafeBuilder represents a safe builder
type SafeBuilder interface {
	Create() SafeBuilder
	WithID(id *uuid.UUID) SafeBuilder
	WithMetaData(met metadata.MetaData) SafeBuilder
	WithToken(tok tokens.Token) SafeBuilder
	WithCipher(cipher cryptography.Cipher) SafeBuilder
	CreatedOn(crOn time.Time) SafeBuilder
	LastUpdatedOn(lstUpOn time.Time) SafeBuilder
	Now() (Safe, error)
}
