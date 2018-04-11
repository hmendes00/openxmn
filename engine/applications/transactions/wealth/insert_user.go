package wealth

import (
	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	concrete_cryptography "github.com/XMNBlockchain/openxmn/engine/infrastructure/cryptography"
	uuid "github.com/satori/go.uuid"
)

// InsertUser represents a save user transaction
type InsertUser struct {
	UserID *uuid.UUID                       `json:"user_id"`
	PubKey *concrete_cryptography.PublicKey `json:"public_key"`
}

// CreateInsertUser creates a new InsertUser instance
func CreateInsertUser(userID *uuid.UUID, pubKey *concrete_cryptography.PublicKey) *InsertUser {
	out := InsertUser{
		UserID: userID,
		PubKey: pubKey,
	}

	return &out
}

// GetUserID returns the userID
func (sav *InsertUser) GetUserID() *uuid.UUID {
	return sav.UserID
}

// GetPublicKey returns the public key
func (sav *InsertUser) GetPublicKey() cryptography.PublicKey {
	return sav.PubKey
}
