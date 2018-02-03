package domain

import (
	"time"

	body "github.com/XMNBlockchain/core/packages/transactions/transactions/domain/body"
	uuid "github.com/satori/go.uuid"
)

// Transaction represents a Transaction
type Transaction interface {
	GetID() *uuid.UUID
	GetBody() body.Body
	GetKarma() int
	CreatedOn() time.Time
}
