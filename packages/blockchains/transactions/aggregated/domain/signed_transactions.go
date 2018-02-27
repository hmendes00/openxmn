package domain

import (
	"time"

	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	uuid "github.com/satori/go.uuid"
)

// SignedTransactions represents the aggregated signed transactions, then signed by a pointer
type SignedTransactions interface {
	GetID() *uuid.UUID
	GetTrs() Transactions
	GetSignature() users.Signature
	CreatedOn() time.Time
}
