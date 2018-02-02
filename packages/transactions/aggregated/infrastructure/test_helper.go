package infrastructure

import (
	"testing"
	"time"

	concrete_signed "github.com/XMNBlockchain/core/packages/transactions/signed/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// CreateTransactionsForTests creates an Transactions instance for tests
func CreateTransactionsForTests(t *testing.T) *Transactions {
	//variables:
	trs := []*concrete_signed.Transaction{
		concrete_signed.CreateTransactionForTests(t),
		concrete_signed.CreateTransactionForTests(t),
	}

	atomicTrs := []*concrete_signed.AtomicTransaction{
		concrete_signed.CreateAtomicTransactionForTests(t),
		concrete_signed.CreateAtomicTransactionForTests(t),
		concrete_signed.CreateAtomicTransactionForTests(t),
	}

	aggregatedTrs := createTransactions(trs, atomicTrs)
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
