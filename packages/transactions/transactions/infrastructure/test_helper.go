package infrastructure

import (
	"math/rand"
	"testing"
	"time"

	concrete_body "github.com/XMNBlockchain/core/packages/transactions/transactions/infrastructure/body"
	uuid "github.com/satori/go.uuid"
)

// CreateTransactionForTests creates a Transaction for tests
func CreateTransactionForTests(t *testing.T) *Transaction {
	//variables:
	id := uuid.NewV4()
	createdOn := time.Now().UTC()
	karma := rand.Int() % 20
	bod := concrete_body.CreateBodyWithUserForTests(t)

	trs := createTransaction(&id, karma, bod, createdOn)
	return trs.(*Transaction)
}
