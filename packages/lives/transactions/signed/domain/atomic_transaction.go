package domain

import (
	"time"

	transactions "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	uuid "github.com/satori/go.uuid"
)

// AtomicTransaction represents multiple transactions that must be executed all at once, signed by a wallet
type AtomicTransaction interface {
	GetID() *uuid.UUID
	GetTrs() []transactions.Transaction
	GetSignature() users.Signature
	CreatedOn() time.Time
}
