package infrastructure

import (
	"time"

	met "github.com/XMNBlockchain/core/packages/lives/metadata/domain"
	uuid "github.com/satori/go.uuid"
)

// MetaData represents a concrete metadata implementation
type MetaData struct {
	ID   *uuid.UUID `json:"id"`
	CrOn time.Time  `json:"created_on"`
}

func createMetaData(id *uuid.UUID, crOn time.Time) met.MetaData {
	out := MetaData{
		ID:   id,
		CrOn: crOn,
	}

	return &out
}

// GetID returns the ID
func (met *MetaData) GetID() *uuid.UUID {
	return met.ID
}

// CreatedOn returns the creation time
func (met *MetaData) CreatedOn() time.Time {
	return met.CrOn
}
