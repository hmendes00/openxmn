package domain

import (
	signed_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	uuid "github.com/satori/go.uuid"
)

// Complaint represents a missing transaction complaint
type Complaint interface {
	GetID() *uuid.UUID
	HasSignedTrs() bool
	GetSignedTrs() signed_transactions.Transaction
	HasAtomicTrs() bool
	GetAtomicTrs() signed_transactions.AtomicTransaction
}
