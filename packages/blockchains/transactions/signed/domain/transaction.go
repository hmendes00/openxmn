package domain

import (
	"time"

	transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	uuid "github.com/satori/go.uuid"
)

// Transaction represents a signed transaction
type Transaction interface {
	GetID() *uuid.UUID
	GetTrs() transactions.Transaction
	GetSignature() users.Signature
	CreatedOn() time.Time
}
