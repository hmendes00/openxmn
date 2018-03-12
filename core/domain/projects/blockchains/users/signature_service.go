package domain

import (
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/users"
)

// SignatureService represents a Signature service
type SignatureService interface {
	Save(dirPath string, sig Signature) (stored_users.Signature, error)
	SaveAll(dirPath string, sigs []Signature) ([]stored_users.Signature, error)
}
