package domain

import (
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/users"
)

// SignaturesService represents a signatures service
type SignaturesService interface {
	Save(dirPath string, sig Signatures) (stored_users.Signatures, error)
}
