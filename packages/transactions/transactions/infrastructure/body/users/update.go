package users

import (
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	"github.com/XMNBlockchain/core/packages/transactions/transactions/domain/body/users"
	uuid "github.com/satori/go.uuid"
)

// Update represents the concrete update pointer transaction
type Update struct {
	ID *uuid.UUID                       `json:"id"`
	PK *concrete_cryptography.PublicKey `json:"public_key"`
}

func createUpdate(id *uuid.UUID, pk *concrete_cryptography.PublicKey) users.Update {
	out := Update{
		ID: id,
		PK: pk,
	}

	return &out
}

// GetID returns the ID of the update wallet transaction
func (up *Update) GetID() *uuid.UUID {
	return up.ID
}

// GetNewPublicKey returns the new PublicKey of the update wallet transaction
func (up *Update) GetNewPublicKey() cryptography.PublicKey {
	return up.PK
}
