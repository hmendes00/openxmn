package objects

import (
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	uuid "github.com/satori/go.uuid"
)

// User represents a user
type User struct {
	ID     *uuid.UUID             `json:"id"`
	PubKey cryptography.PublicKey `json:"public_key"`
	Wal    map[string]*Wallet     `json:"wallet"`
}
