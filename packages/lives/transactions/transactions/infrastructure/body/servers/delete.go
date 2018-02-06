package servers

import (
	servers "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/servers"
	uuid "github.com/satori/go.uuid"
)

// Delete represents the concrete delete pointer transaction
type Delete struct {
	ID *uuid.UUID `json:"id"`
}

func createDelete(id *uuid.UUID) servers.Delete {
	out := Delete{
		ID: id,
	}

	return &out
}

// GetID returns the ID of the pointer
func (del *Delete) GetID() *uuid.UUID {
	return del.ID
}
