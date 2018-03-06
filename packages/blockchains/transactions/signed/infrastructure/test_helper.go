package infrastructure

import (
	"strconv"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	concrete_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/blockchains/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// CreateTransactionForTests creates a Transaction for tests
func CreateTransactionForTests(t *testing.T) *Transaction {
	//variables:
	id := uuid.NewV4()
	trs := concrete_transactions.CreateTransactionForTests(t)
	sig := concrete_users.CreateSignatureForTests(t)
	createdOn := time.Now().UTC()

	blocks := [][]byte{
		id.Bytes(),
		trs.GetMetaData().GetHashTree().GetHash().Get(),
		[]byte(sig.GetSig().String()),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
	}
	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()

	sigTrs := createTransaction(met.(*concrete_metadata.MetaData), trs, sig)
	return sigTrs.(*Transaction)
}

// CreateTransactionsForTests creates a Transactions for tests
func CreateTransactionsForTests(t *testing.T) *Transactions {
	//variables:
	id := uuid.NewV4()
	createdOn := time.Now().UTC()
	trs := []*Transaction{
		CreateTransactionForTests(t),
		CreateTransactionForTests(t),
		CreateTransactionForTests(t),
		CreateTransactionForTests(t),
	}

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
	}

	for _, oneTrs := range trs {
		blocks = append(blocks, oneTrs.GetMetaData().GetHashTree().GetHash().Get())
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()

	sigTrs := createTransactions(met.(*concrete_metadata.MetaData), trs)
	return sigTrs.(*Transactions)
}

// CreateAtomicTransactionForTests creates an AtomicTransaction for tests
func CreateAtomicTransactionForTests(t *testing.T) *AtomicTransaction {
	//variables:
	id := uuid.NewV4()
	trs := concrete_transactions.CreateTransactionsForTests(t)
	sig := concrete_users.CreateSignatureForTests(t)
	createdOn := time.Now().UTC()

	blocks := [][]byte{
		id.Bytes(),
		trs.GetMetaData().GetHashTree().GetHash().Get(),
		[]byte(sig.GetSig().String()),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()

	atomicTrs := createAtomicTransaction(met.(*concrete_metadata.MetaData), trs, sig)
	return atomicTrs.(*AtomicTransaction)
}

// CreateAtomicTransactionsForTests creates an AtomicTransaction for tests
func CreateAtomicTransactionsForTests(t *testing.T) *AtomicTransactions {
	//variables:
	id := uuid.NewV4()
	createdOn := time.Now().UTC()
	atomicTrs := []*AtomicTransaction{
		CreateAtomicTransactionForTests(t),
		CreateAtomicTransactionForTests(t),
		CreateAtomicTransactionForTests(t),
		CreateAtomicTransactionForTests(t),
		CreateAtomicTransactionForTests(t),
	}

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
	}

	for _, oneAtomicTrs := range atomicTrs {
		blocks = append(blocks, oneAtomicTrs.GetMetaData().GetHashTree().GetHash().Get())
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()

	out := createAtomicTransactions(met.(*concrete_metadata.MetaData), atomicTrs)
	return out.(*AtomicTransactions)
}
