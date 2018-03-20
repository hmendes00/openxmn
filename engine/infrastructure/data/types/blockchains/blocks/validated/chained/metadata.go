package chained

import (
	"time"

	chained "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/blocks/validated/chained"
	hashtrees "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/hashtrees"
	concrete_hashtrees "github.com/XMNBlockchain/exmachina-network/engine/infrastructure/data/types/hashtrees"
	uuid "github.com/satori/go.uuid"
)

// MetaData represents a concrete metadata implementation
type MetaData struct {
	ID     *uuid.UUID                   `json:"id"`
	HT     *concrete_hashtrees.HashTree `json:"hashtree"`
	PrevID *uuid.UUID                   `json:"previous_id"`
	CrOn   time.Time                    `json:"created_on"`
}

func createMetaData(id *uuid.UUID, ht *concrete_hashtrees.HashTree, prevID *uuid.UUID, createdOn time.Time) chained.MetaData {
	out := MetaData{
		ID:     id,
		HT:     ht,
		PrevID: prevID,
		CrOn:   createdOn,
	}

	return &out
}

// GetID returns the ID
func (met *MetaData) GetID() *uuid.UUID {
	return met.ID
}

// GetHashTree returns the HashTree
func (met *MetaData) GetHashTree() hashtrees.HashTree {
	return met.HT
}

// GetPreviousID returns the previous ID
func (met *MetaData) GetPreviousID() *uuid.UUID {
	return met.PrevID
}

// CreatedOn returns the creation time
func (met *MetaData) CreatedOn() time.Time {
	return met.CrOn
}
