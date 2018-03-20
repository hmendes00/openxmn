package transactions

import (
	uuid "github.com/satori/go.uuid"
)

// DeleteServer represents a delete server transaction
type DeleteServer struct {
	ServerID *uuid.UUID `json:"server_id"`
}
