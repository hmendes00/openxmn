package rsa

import (
	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
)

// PrivateKeyBuilderFactory represents a concrete PrivateKeyBuilder factory
type PrivateKeyBuilderFactory struct {
}

// CreatePrivateKeyBuilderFactory creates a new PrivateKeyBuilderFactory instance
func CreatePrivateKeyBuilderFactory() cryptography.PrivateKeyBuilderFactory {
	out := PrivateKeyBuilderFactory{}

	return &out
}

// Create creates a new PrivateKeyBuilder instance
func (fac *PrivateKeyBuilderFactory) Create() cryptography.PrivateKeyBuilder {
	out := createPrivateKeyBuilder()
	return out
}
