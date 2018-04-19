package servers

import (
	uuid "github.com/satori/go.uuid"
)

// DeleteServer represents a delete server transaction
type DeleteServer struct {
	ServerID *uuid.UUID `json:"server_id"`
}

// CreateDeleteServer creates a new DeleteServer instance
func CreateDeleteServer(serverID *uuid.UUID) *DeleteServer {
	out := DeleteServer{
		ServerID: serverID,
	}

	return &out
}

// GetServerID returns the server ID
func (del *DeleteServer) GetServerID() *uuid.UUID {
	return del.ServerID
}
