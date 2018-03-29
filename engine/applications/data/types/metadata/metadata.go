package metadata

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// MetaData represents the metadata of a type
type MetaData struct {
	ID       *uuid.UUID `json:"id"`
	CrOn     time.Time  `json:"created_on"`
	LastUpOn time.Time  `json:"last_updated_on"`
}

// CreateMetaData creates a metadata instance
func CreateMetaData(id *uuid.UUID, crOn time.Time, lastUpOn time.Time) *MetaData {
	out := MetaData{
		ID:       id,
		CrOn:     crOn,
		LastUpOn: lastUpOn,
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

// LastUpdatedOn returns the last updated on time
func (met *MetaData) LastUpdatedOn() time.Time {
	return met.LastUpOn
}
