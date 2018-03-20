package cryptography

import (
	"crypto/rsa"
	"net/url"
)

// SignatureBuilder represents a Signature builder
type SignatureBuilder interface {
	Create() SignatureBuilder
	WithPrivateKey(pk *rsa.PrivateKey) SignatureBuilder
	WithPublicKey(pub *rsa.PublicKey) SignatureBuilder
	WithEncodedPublicKey(encodedPub string) SignatureBuilder
	WithData(data []byte) SignatureBuilder
	WithURLValues(urlValues url.Values) SignatureBuilder
	WithInterface(v interface{}) SignatureBuilder
	WithSignature(sig []byte) SignatureBuilder
	WithEncodedSignature(encodedSig string) SignatureBuilder
	Now() (Signature, error)
}
