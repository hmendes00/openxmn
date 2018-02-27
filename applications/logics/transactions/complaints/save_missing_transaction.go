package complaints

import (
	signed_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// SaveMissingTransaction represents a save missing transaction complaint
type SaveMissingTransaction struct {
	ID              *uuid.UUID                             `json:"id"`
	SignedTrs       *signed_transactions.Transaction       `json:"signed_transaction"`
	AtomicSignedTrs *signed_transactions.AtomicTransaction `json:"atomic_signed_transactions"`
}
