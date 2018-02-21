package domain

import (
	users "github.com/XMNBlockchain/core/packages/lives/users/domain"
)

// Signature represents an API signature
type Signature interface {
	HasSignature() bool
	GetSignature() users.Signature
}
