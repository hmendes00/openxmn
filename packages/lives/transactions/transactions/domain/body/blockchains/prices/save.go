package prices

import uuid "github.com/satori/go.uuid"

// Save represents price related to a blockchain
type Save interface {
	GetBlockchainID() *uuid.UUID
	GetTransaction() float64
	GetIncomingBandwidthPerMb() float64
	GetOutgoingBandwidthPerMb() float64
	GetStoragePerMbPerHour() float64
}
