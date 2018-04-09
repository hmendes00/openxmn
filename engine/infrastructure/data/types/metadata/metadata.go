package metadata

import (
	"time"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	uuid "github.com/satori/go.uuid"
)

// MetaData represents metadata
type MetaData struct {
	ID      *uuid.UUID `json:"metadata"`
	CrOn    time.Time  `json:"created_on"`
	LstUpOn time.Time  `json:"last_updated_on"`
}

func createMetaData(id *uuid.UUID, crOn time.Time, lstUpOn time.Time) metadata.MetaData {
	out := MetaData{
		ID:      id,
		CrOn:    crOn,
		LstUpOn: lstUpOn,
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
	return met.LstUpOn
}
