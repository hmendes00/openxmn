package infrastructure

import (
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	user "github.com/XMNBlockchain/core/packages/users/domain"
	uuid "github.com/satori/go.uuid"
)

// User represents the concrete user
type User struct {
	ID *uuid.UUID                       `json:"id"`
	PK *concrete_cryptography.PublicKey `json:"public_key"`
}

func createUser(id *uuid.UUID, pub *concrete_cryptography.PublicKey) user.User {
	out := User{
		ID: id,
		PK: pub,
	}

	return &out
}

// GetID returns the user ID
func (user *User) GetID() *uuid.UUID {
	return user.ID
}

// GetPublicKey returns the PublicKey
func (user *User) GetPublicKey() cryptography.PublicKey {
	return user.PK
}
