package transactions

import uuid "github.com/satori/go.uuid"

// ServerRunInstance represents a transaction used to run an instance
type ServerRunInstance struct {
	ID         *uuid.UUID `json:"id"`
	InstanceID *uuid.UUID `json:"instance_id"`
	ServerID   *uuid.UUID `json:"server_id"`
	StakeID    *uuid.UUID `json:"stake_id"`
}
