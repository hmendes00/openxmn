package infrastructure

import (
	"strconv"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	concrete_signed "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/blockchains/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// CreateTransactionsForTests creates an Transactions instance for tests
func CreateTransactionsForTests(t *testing.T) *Transactions {
	//variables:
	id := uuid.NewV4()
	trs := concrete_signed.CreateTransactionsForTests(t)
	atomicTrs := concrete_signed.CreateAtomicTransactionsForTests(t)
	createdOn := time.Now().UTC()

	htBlocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
		trs.GetMetaData().GetHashTree().GetHash().Get(),
		atomicTrs.GetMetaData().GetHashTree().GetHash().Get(),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(htBlocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()

	aggregatedTrs := createTransactions(met.(*concrete_metadata.MetaData), trs, atomicTrs)
	return aggregatedTrs.(*Transactions)
}

// CreateSignedTransactionsForTests creates a SignedTransactions instance for tests
func CreateSignedTransactionsForTests(t *testing.T) *SignedTransactions {
	//variables:
	id := uuid.NewV4()
	trs := CreateTransactionsForTests(t)
	sig := concrete_users.CreateSignatureForTests(t)
	cr := time.Now().UTC()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(cr.UnixNano()))),
		trs.GetMetaData().GetHashTree().GetHash().Get(),
		[]byte(sig.GetSig().String()),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(cr).Now()

	sigTrs := createSignedTransactions(met.(*concrete_metadata.MetaData), trs, sig)
	return sigTrs.(*SignedTransactions)
}

// CreateAggregatedSignedTransactionsForTests creates an AggregatedSignedTransactions instance for tests
func CreateAggregatedSignedTransactionsForTests(t *testing.T) *AggregatedSignedTransactions {
	//variables:
	id := uuid.NewV4()
	cr := time.Now().UTC()
	trs := []*SignedTransactions{
		CreateSignedTransactionsForTests(t),
		CreateSignedTransactionsForTests(t),
		CreateSignedTransactionsForTests(t),
	}

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(cr.UnixNano()))),
	}

	for _, oneTrs := range trs {
		blocks = append(blocks, oneTrs.GetMetaData().GetHashTree().GetHash().Get())
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(cr).Now()

	aggrSignedTrs := createAggregatedSignedTransactions(met.(*concrete_metadata.MetaData), trs)
	return aggrSignedTrs.(*AggregatedSignedTransactions)
}
