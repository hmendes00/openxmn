package domain

import (
	stored_file "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// SignatureService represents a Signature service
type SignatureService interface {
	Save(dirPath string, sig Signature) (stored_file.File, error)
	SaveAll(dirPath string, sigs []Signature) ([]stored_file.File, error)
}
