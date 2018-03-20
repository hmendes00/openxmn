package cryptography

// Cipher represents encrypted data
type Cipher interface {
	GetText() []byte
	GetLabel() []byte
	GetSignature() Signature
	Decipher(pk PrivateKey) ([]byte, error)
}
