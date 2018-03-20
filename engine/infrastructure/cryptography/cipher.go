package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"hash"

	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
)

// Cipher represents a concrete cipher implementation
type Cipher struct {
	h     hash.Hash
	text  []byte
	label []byte
	sig   *Signature
}

type jsonifyCipher struct {
	Hash  string     `json:"hash"`
	Text  string     `json:"data"`
	Label string     `json:"data"`
	Sig   *Signature `json:"signature"`
}

func createCipher(h hash.Hash, text []byte, label []byte, sig *Signature) cryptography.Cipher {
	out := Cipher{
		h:     h,
		text:  text,
		label: label,
		sig:   sig,
	}

	return &out
}

// GetText returns the cipher text
func (ci *Cipher) GetText() []byte {
	return ci.text
}

// GetLabel returns the label
func (ci *Cipher) GetLabel() []byte {
	return ci.label
}

// GetSignature returns the signature
func (ci *Cipher) GetSignature() cryptography.Signature {
	return ci.sig
}

// Decipher convert a cipher to clear data
func (ci *Cipher) Decipher(pk cryptography.PrivateKey) ([]byte, error) {
	//make sure the given PK contains a public key that is the same as the signature:
	if !ci.sig.GetPublicKey().Compare(pk.GetPublicKey()) {
		return nil, errors.New("the given PrivateKey does not contain a PublicKey that was used to sign the cipher")
	}

	//decrypt the cipher:
	txt, txtErr := rsa.DecryptOAEP(ci.h, rand.Reader, pk.GetKey(), ci.text, ci.label)
	if txtErr != nil {
		return nil, txtErr
	}

	return txt, nil
}
