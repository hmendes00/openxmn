package users

import (
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	uuid "github.com/satori/go.uuid"
)

// CreateBuilder represents the builder of a create user transaction
type CreateBuilder interface {
	Create() CreateBuilder
	WithID(id *uuid.UUID) CreateBuilder
	WithPublicKey(pk cryptography.PublicKey) CreateBuilder
	Now() (Create, error)
}
