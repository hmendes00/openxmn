package domain

import (
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	uuid "github.com/satori/go.uuid"
)

// AtomicTransaction represents multiple transactions that must be executed all at once, signed by a wallet
type AtomicTransaction interface {
	GetID() *uuid.UUID
	GetHashTree() hashtrees.HashTree
	GetTrs() []transactions.Transaction
	GetSignature() users.Signature
	CreatedOn() time.Time
}
