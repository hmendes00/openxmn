package infrastructure

import (
	"time"

	chained "github.com/XMNBlockchain/core/packages/blockchains/blocks/chained/domain"
	uuid "github.com/satori/go.uuid"
)

// MetaData represents a concrete metadata implementation
type MetaData struct {
	ID        *uuid.UUID `json:"id"`
	Index     int        `json:"index"`
	PrevIndex int        `json:"previous_index"`
	CrOn      time.Time  `json:"created_on"`
}

func createMetaData(id *uuid.UUID, index int, prevIndex int, createdOn time.Time) chained.MetaData {
	out := MetaData{
		ID:        id,
		Index:     index,
		PrevIndex: prevIndex,
		CrOn:      createdOn,
	}

	return &out
}

// GetID returns the ID
func (met *MetaData) GetID() *uuid.UUID {
	return met.ID
}

// GetIndex returns the index
func (met *MetaData) GetIndex() int {
	return met.Index
}

// GetPreviousIndex returns the previous index
func (met *MetaData) GetPreviousIndex() int {
	return met.PrevIndex
}

// CreatedOn returns the creation time
func (met *MetaData) CreatedOn() time.Time {
	return met.CrOn
}
