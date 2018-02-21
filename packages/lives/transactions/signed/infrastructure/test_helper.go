package infrastructure

import (
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/infrastructure"
	concrete_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/lives/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// CreateTransactionForTests creates a Transaction for tests
func CreateTransactionForTests(t *testing.T) *Transaction {
	//variables:
	id := uuid.NewV4()
	trs := concrete_transactions.CreateTransactionForTests(t)
	sig := concrete_users.CreateSignatureForTests(t)
	createdOn := time.Now().UTC()

	sigTrs := createTransaction(&id, trs, sig, createdOn)
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

	htBlocks := [][]byte{}
	for _, onTrs := range trs {
		htBlocks = append(htBlocks, onTrs.GetID().Bytes())
	}

	//create hashtree:
	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(htBlocks).Now()

	atomicTrs := createAtomicTransaction(&id, ht.(*concrete_hashtrees.HashTree), trs, sig, createdOn)
	return atomicTrs.(*AtomicTransaction)
}
