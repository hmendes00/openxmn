package domain

import (
	"time"

	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	uuid "github.com/satori/go.uuid"
)

// SignatureBuilder represents the Signature builder
type SignatureBuilder interface {
	Create() SignatureBuilder
	WithID(id *uuid.UUID) SignatureBuilder
	WithMetaData(met metadata.MetaData) SignatureBuilder
	WithPrivateKey(pk cryptography.PrivateKey) SignatureBuilder
	WithInterface(v interface{}) SignatureBuilder
	WithSignature(sig cryptography.Signature) SignatureBuilder
	WithUser(usr User) SignatureBuilder
	CreatedOn(crOn time.Time) SignatureBuilder
	Now() (Signature, error)
}
