package transactions

import uuid "github.com/satori/go.uuid"

// SaveApplication represents a save application transactiomn
type SaveApplication struct {
	ApplicationID *uuid.UUID `json:"application_id"`
	Version       string     `json:"version"`
	Name          string     `json:"name"`
	BlockchainID  *uuid.UUID `json:"blockchain_id"`
	DockerFile    string     `json:"docker_file"`
}
