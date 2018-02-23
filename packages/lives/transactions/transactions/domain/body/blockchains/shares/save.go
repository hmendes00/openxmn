package shares

import uuid "github.com/satori/go.uuid"

// Save represents the shares split for the transactions
type Save interface {
	GetBlockchainID() *uuid.UUID
	GetTransaction() float64
	GetLeader() float64
	GetVerifier() float64
}
