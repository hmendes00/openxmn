package domain

import (
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/users"
)

// SignaturesService represents a signatures service
type SignaturesService interface {
	Save(dirPath string, sig Signatures) (stored_users.Signatures, error)
}
