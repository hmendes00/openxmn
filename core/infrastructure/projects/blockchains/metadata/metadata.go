package metadata

import (
	"time"

	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/hashtrees"
	met "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
	concrete_hashtrees "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/hashtrees"
	uuid "github.com/satori/go.uuid"
)

// MetaData represents a concrete metadata implementation
type MetaData struct {
	ID   *uuid.UUID                   `json:"id"`
	HT   *concrete_hashtrees.HashTree `json:"hashtree"`
	CrOn time.Time                    `json:"created_on"`
}

func createMetaData(id *uuid.UUID, ht *concrete_hashtrees.HashTree, crOn time.Time) met.MetaData {
	out := MetaData{
		ID:   id,
		HT:   ht,
		CrOn: crOn,
	}

	return &out
}

// GetID returns the ID
func (met *MetaData) GetID() *uuid.UUID {
	return met.ID
}

// GetHashTree returns the hashtree
func (met *MetaData) GetHashTree() hashtrees.HashTree {
	return met.HT
}

// CreatedOn returns the creation time
func (met *MetaData) CreatedOn() time.Time {
	return met.CrOn
}
