package infrastructure

import (
	"testing"
	"time"

	concrete_body "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure/body"
	uuid "github.com/satori/go.uuid"
)

// CreateTransactionForTests creates a Transaction for tests
func CreateTransactionForTests(t *testing.T) *Transaction {
	//variables:
	id := uuid.NewV4()
	createdOn := time.Now().UTC()
	bod := concrete_body.CreateBodyWithUserForTests(t)

	trs := createTransaction(&id, bod, createdOn)
	return trs.(*Transaction)
}
