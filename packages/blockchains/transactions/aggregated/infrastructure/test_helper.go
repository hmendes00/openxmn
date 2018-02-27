package infrastructure

import (
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_signed "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/blockchains/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// CreateTransactionsForTests creates an Transactions instance for tests
func CreateTransactionsForTests(t *testing.T) *Transactions {
	//variables:
	id := uuid.NewV4()
	trs := []*concrete_signed.Transaction{
		concrete_signed.CreateTransactionForTests(t),
		concrete_signed.CreateTransactionForTests(t),
	}

	atomicTrs := []*concrete_signed.AtomicTransaction{
		concrete_signed.CreateAtomicTransactionForTests(t),
		concrete_signed.CreateAtomicTransactionForTests(t),
		concrete_signed.CreateAtomicTransactionForTests(t),
	}

	createdOn := time.Now().UTC()

	htBlocks := [][]byte{}
	for _, oneTrs := range trs {
		htBlocks = append(htBlocks, oneTrs.GetID().Bytes())
	}

	for _, oneAtomicTrs := range atomicTrs {
		htBlocks = append(htBlocks, oneAtomicTrs.GetID().Bytes())
	}

	//create hashtree:
	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(htBlocks).Now()

	aggregatedTrs := createTransactions(&id, ht.(*concrete_hashtrees.HashTree), trs, atomicTrs, createdOn)
	return aggregatedTrs.(*Transactions)
}

// CreateSignedTransactionsForTests creates a SignedTransactions instance for tests
func CreateSignedTransactionsForTests(t *testing.T) *SignedTransactions {
	//variables:
	id := uuid.NewV4()
	trs := CreateTransactionsForTests(t)
	sig := concrete_users.CreateSignatureForTests(t)
	cr := time.Now().UTC()

	sigTrs := createSignedTransactions(&id, trs, sig, cr)
	return sigTrs.(*SignedTransactions)
}
