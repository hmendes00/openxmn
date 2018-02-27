package infrastructure

import (
	"crypto/sha256"
	"encoding/hex"

	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	user "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
)

// Signature represents the concrete user signature
type Signature struct {
	Sig *concrete_cryptography.Signature `json:"signature"`
	Usr *User                            `json:"user"`
}

func createSignature(sig *concrete_cryptography.Signature, usr *User) user.Signature {
	out := Signature{
		Sig: sig,
		Usr: usr,
	}

	return &out
}

// GetKey returns the key
func (sig *Signature) GetKey() string {
	str := sig.GetSig().String()
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// GetSig returns the signature
func (sig *Signature) GetSig() cryptography.Signature {
	return sig.Sig
}

// GetUser returns the user
func (sig *Signature) GetUser() user.User {
	return sig.Usr
}
