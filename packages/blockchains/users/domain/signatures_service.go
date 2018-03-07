package domain

import (
	stored_users "github.com/XMNBlockchain/core/packages/storages/users/domain"
)

// SignaturesService represents a signatures service
type SignaturesService interface {
	Save(dirPath string, sig Signatures) (stored_users.Signatures, error)
}
