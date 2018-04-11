package wealth

import (
	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	concrete_cryptography "github.com/XMNBlockchain/openxmn/engine/infrastructure/cryptography"
)

// UpdateUser represents a save user transaction
type UpdateUser struct {
	PubKey *concrete_cryptography.PublicKey `json:"public_key"`
}

// CreateUpdateUser creates a new UpdateUser instance
func CreateUpdateUser(pubKey *concrete_cryptography.PublicKey) *UpdateUser {
	out := UpdateUser{
		PubKey: pubKey,
	}

	return &out
}

// GetPublicKey returns the public key
func (sav *UpdateUser) GetPublicKey() cryptography.PublicKey {
	return sav.PubKey
}
