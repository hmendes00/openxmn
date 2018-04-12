package databases

import (
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/files"
	safes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/safes"
	uuid "github.com/satori/go.uuid"
)

// Safe represents a safe database
type Safe struct {
	dirPath    string
	repository safes.SafeRepository
}

// CreateSafe creates a new Safe instance
func CreateSafe(dirPath string, repository safes.SafeRepository) *Safe {
	out := Safe{
		dirPath:    dirPath,
		repository: repository,
	}

	return &out
}

// RetrieveByID retrieves the safe by ID
func (saf *Safe) RetrieveByID(id *uuid.UUID) (safes.Safe, error) {
	return nil, nil
}

// Update updates a safe
func (saf *Safe) Update(old safes.Safe, new safes.Safe) (files.File, files.File, error) {
	return nil, nil, nil
}

// Delete deletes a safe
func (saf *Safe) Delete(safe safes.Safe) (files.File, error) {
	return nil, nil
}
