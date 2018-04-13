package users

import (
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/users"
)

// SignatureBuilderFactory represents a concrete SignatureBuilderFactory implementation
type SignatureBuilderFactory struct {
}

// CreateSignatureBuilderFactory creates a new SignatureBuilderFactory instance
func CreateSignatureBuilderFactory() stored_users.SignatureBuilderFactory {
	out := SignatureBuilderFactory{}
	return &out
}

// Create creates a new SignatureBuilder instance
func (fac *SignatureBuilderFactory) Create() stored_users.SignatureBuilder {
	out := createSignatureBuilder()
	return out
}
