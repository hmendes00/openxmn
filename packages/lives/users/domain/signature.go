package domain

import (
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
)

// Signature represents the Signature of a User
type Signature interface {
	GetKey() string
	GetSig() cryptography.Signature
	GetUser() User
}
