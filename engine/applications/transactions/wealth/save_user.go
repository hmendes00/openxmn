package wealth

import (
	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	concrete_cryptography "github.com/XMNBlockchain/openxmn/engine/infrastructure/cryptography"
	uuid "github.com/satori/go.uuid"
)

// SaveUser represents a save user transaction
type SaveUser struct {
	UserID *uuid.UUID                       `json:"user_id"`
	PubKey *concrete_cryptography.PublicKey `json:"public_key"`
}

// CreateSaveUser creates a new SaveUser instance
func CreateSaveUser(userID *uuid.UUID, pubKey *concrete_cryptography.PublicKey) *SaveUser {
	out := SaveUser{
		UserID: userID,
		PubKey: pubKey,
	}

	return &out
}

// GetUserID returns the userID
func (sav *SaveUser) GetUserID() *uuid.UUID {
	return sav.UserID
}

// GetPublicKey returns the public key
func (sav *SaveUser) GetPublicKey() cryptography.PublicKey {
	return sav.PubKey
}
