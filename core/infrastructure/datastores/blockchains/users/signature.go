package users

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	cryptography "github.com/XMNBlockchain/exmachina-network/core/domain/cryptography"
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/metadata"
	user "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/users"
	concrete_cryptography "github.com/XMNBlockchain/exmachina-network/core/infrastructure/cryptography/rsa"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/metadata"
)

// Signature represents the concrete user signature
type Signature struct {
	Met *concrete_metadata.MetaData      `json:"metadata"`
	Sig *concrete_cryptography.Signature `json:"signature"`
	Usr *User                            `json:"user"`
}

func createSignature(met *concrete_metadata.MetaData, sig *concrete_cryptography.Signature, usr *User) (user.Signature, error) {

	//make sure the public key matches:
	sigPubKey := sig.GetPublicKey()
	if !usr.GetPublicKey().Compare(sigPubKey) {
		return nil, errors.New("the PublicKey of the Signature does not match the PublicKey of the User instance")
	}

	out := Signature{
		Met: met,
		Sig: sig,
		Usr: usr,
	}

	return &out, nil
}

// GetMetaData returns the MetaData
func (sig *Signature) GetMetaData() metadata.MetaData {
	return sig.Met
}

// GetSignature returns the signature
func (sig *Signature) GetSignature() cryptography.Signature {
	return sig.Sig
}

// GetUser returns the user
func (sig *Signature) GetUser() user.User {
	return sig.Usr
}

// String returns an encoded signature as string
func (sig *Signature) String() string {
	js, jsErr := json.Marshal(sig)
	if jsErr != nil {
		return ""
	}

	encoded := base64.URLEncoding.EncodeToString(js)
	return encoded
}
