package custom

import (
	custom "github.com/XMNBlockchain/core/packages/transactions/transactions/domain/body/custom"
	uuid "github.com/satori/go.uuid"
)

// Create represents the concrete create custom transaction
type Create struct {
	ID *uuid.UUID `json:"id"`
	JS []byte     `json:"json"`
}

func createCreate(id *uuid.UUID, js []byte) custom.Create {
	out := Create{
		ID: id,
		JS: js,
	}

	return &out
}

// GetID returns the ID
func (cr *Create) GetID() *uuid.UUID {
	return cr.ID
}

// GetJSON returns the JSON data
func (cr *Create) GetJSON() []byte {
	return cr.JS
}
