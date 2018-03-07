package infrastructure

import (
	stored_users "github.com/XMNBlockchain/core/packages/storages/users/domain"
)

// SignaturesBuilderFactory represents a concrete SignaturesBuilderFactory implementation
type SignaturesBuilderFactory struct {
}

// CreateSignaturesBuilderFactory creates a new SignaturesBuilderFactory instance
func CreateSignaturesBuilderFactory() stored_users.SignaturesBuilderFactory {
	out := SignaturesBuilderFactory{}
	return &out
}

// Create creates a new SignaturesBuilder instance
func (fac *SignaturesBuilderFactory) Create() stored_users.SignaturesBuilder {
	out := createSignaturesBuilder()
	return out
}
