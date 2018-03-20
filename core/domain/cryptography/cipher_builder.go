package cryptography

import "hash"

// CipherBuilder represents a cipher builder
type CipherBuilder interface {
	Create() CipherBuilder
	WithPublicKey(pubKey PublicKey) CipherBuilder
	WithText(txt []byte) CipherBuilder
	WithHash(h hash.Hash) CipherBuilder
	WithCipherText(cipherText []byte) CipherBuilder
	WithLabel(label []byte) CipherBuilder
	WithSignature(sig Signature) CipherBuilder
	Now() (Cipher, error)
}
