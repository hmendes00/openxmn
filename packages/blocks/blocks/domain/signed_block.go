package domain

import (
	users "github.com/XMNBlockchain/core/packages/users/domain"
)

// SignedBlock represents a SignedBlock instance
type SignedBlock interface {
	GetBlock() Block
	GetSignature() users.Signature
}
