package domain

import (
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	uuid "github.com/satori/go.uuid"
)

// SignatureBuilder represents the Signature builder
type SignatureBuilder interface {
	Create() SignatureBuilder
	WithPrivateKey(pk cryptography.PrivateKey) SignatureBuilder
	WithInterface(v interface{}) SignatureBuilder
	WithEncodedSignature(encodedSig string) SignatureBuilder
	WithSignature(sig cryptography.Signature) SignatureBuilder
	WithUserID(userID *uuid.UUID) SignatureBuilder
	Now() (Signature, error)
}
