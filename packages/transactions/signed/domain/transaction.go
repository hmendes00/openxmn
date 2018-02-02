package domain

import (
	transactions "github.com/XMNBlockchain/core/packages/transactions/transactions/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	uuid "github.com/satori/go.uuid"
)

// Transaction represents a signed transaction
type Transaction interface {
	GetID() *uuid.UUID
	GetTrs() transactions.Transaction
	GetSignature() users.Signature
}
