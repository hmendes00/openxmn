package databases

import (
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/files"
	stakes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/stakes"
	uuid "github.com/satori/go.uuid"
)

// Stake represents a stake database
type Stake struct {
	dirPath    string
	repository stakes.StakeRepository
}

// CreateStake creates a new Stake instance
func CreateStake(dirPath string, repository stakes.StakeRepository) *Stake {
	out := Stake{
		dirPath:    dirPath,
		repository: repository,
	}

	return &out
}

// RetrieveByID retrieves a Stake by ID
func (db *Stake) RetrieveByID(id *uuid.UUID) (stakes.Stake, error) {
	return nil, nil
}

// Insert insert a new Stake
func (db *Stake) Insert(stk stakes.Stake) (files.File, error) {
	return nil, nil
}

// Delete deletes a stake
func (db *Stake) Delete(stk stakes.Stake) (files.File, error) {
	return nil, nil
}
