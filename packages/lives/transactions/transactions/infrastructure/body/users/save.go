package users

import (
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	users "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/users"
	uuid "github.com/satori/go.uuid"
)

// Save represents the concrete create user transaction
type Save struct {
	ID *uuid.UUID                       `json:"id"`
	PK *concrete_cryptography.PublicKey `json:"public_key"`
}

func createSave(id *uuid.UUID, pk *concrete_cryptography.PublicKey) users.Save {
	out := Save{
		ID: id,
		PK: pk,
	}

	return &out
}

// GetID returns the ID of the create user transaction
func (cr *Save) GetID() *uuid.UUID {
	return cr.ID
}

// GetPublicKey returns the PublicKey of the create user transaction
func (cr *Save) GetPublicKey() cryptography.PublicKey {
	return cr.PK
}
