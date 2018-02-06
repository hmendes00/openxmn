package infrastructure

import (
	"testing"
	"time"

	concrete_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// CreateTransactionForTests creates a Transaction for tests
func CreateTransactionForTests(t *testing.T) *Transaction {
	//variables:
	id := uuid.NewV4()
	trs := concrete_transactions.CreateTransactionForTests(t)
	sig := concrete_users.CreateSignatureForTests(t)

	sigTrs := createTransaction(&id, trs, sig)
	return sigTrs.(*Transaction)
}

// CreateAtomicTransactionForTests creates an AtomicTransaction for tests
func CreateAtomicTransactionForTests(t *testing.T) *AtomicTransaction {
	//variables:
	id := uuid.NewV4()
	trs := []*concrete_transactions.Transaction{
		concrete_transactions.CreateTransactionForTests(t),
		concrete_transactions.CreateTransactionForTests(t),
	}
	sig := concrete_users.CreateSignatureForTests(t)
	createdOn := time.Unix(1515467177, 0)

	atomicTrs := createAtomicTransaction(&id, trs, sig, createdOn)
	return atomicTrs.(*AtomicTransaction)
}
