package users

import (
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	uuid "github.com/satori/go.uuid"
)

// Save represents a save user transaction
type Save struct {
	ID     *uuid.UUID              `json:"id"`
	PubKey *cryptography.PublicKey `json:"public_key"`
}

// CreateSave creates a new Save instance
func CreateSave(id *uuid.UUID, pubKey *cryptography.PublicKey) *Save {
	out := Save{
		ID:     id,
		PubKey: pubKey,
	}

	return &out
}

// GetID returns the ID
func (sav *Save) GetID() *uuid.UUID {
	return sav.ID
}

// GetPublicKey returns the public key
func (sav *Save) GetPublicKey() *cryptography.PublicKey {
	return sav.PubKey
}
