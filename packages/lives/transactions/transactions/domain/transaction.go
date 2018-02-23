package domain

import (
	"time"

	body "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body"
	uuid "github.com/satori/go.uuid"
)

// Transaction represents a Transaction
type Transaction interface {
	GetID() *uuid.UUID
	GetBody() body.Body
	CreatedOn() time.Time
}
