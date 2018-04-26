package delete

import uuid "github.com/satori/go.uuid"

// Server represents a delete server transaction
type Server struct {
	ServerID *uuid.UUID `json:"server_id"`
}
