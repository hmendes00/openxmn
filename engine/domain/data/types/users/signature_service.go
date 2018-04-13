package domain

import (
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/users"
)

// SignatureService represents a Signature service
type SignatureService interface {
	Save(dirPath string, sig Signature) (stored_users.Signature, error)
	SaveAll(dirPath string, sigs []Signature) ([]stored_users.Signature, error)
}
