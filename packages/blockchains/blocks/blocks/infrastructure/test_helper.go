package infrastructure

import (
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/blockchains/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// CreateBlockForTests creates a Block for tests
func CreateBlockForTests(t *testing.T) *Block {
	//variables:
	id := uuid.NewV4()
	crOn := time.Now().UTC()
	trs := []*concrete_aggregated.SignedTransactions{
		concrete_aggregated.CreateSignedTransactionsForTests(t),
		concrete_aggregated.CreateSignedTransactionsForTests(t),
		concrete_aggregated.CreateSignedTransactionsForTests(t),
	}

	htBlocks := [][]byte{}
	for _, oneTrs := range trs {
		htBlocks = append(htBlocks, oneTrs.GetID().Bytes())
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(htBlocks).Now()

	blk := createBlock(&id, ht.(*concrete_hashtrees.HashTree), trs, crOn)
	return blk.(*Block)
}

// CreateSignedBlockForTests creates a SignedBlock for tests
func CreateSignedBlockForTests(t *testing.T) *SignedBlock {
	//variables:
	id := uuid.NewV4()
	blk := CreateBlockForTests(t)
	sig := concrete_users.CreateSignatureForTests(t)
	crOn := time.Now().UTC()

	signedBlk := createSignedBlock(&id, blk, sig, crOn)
	return signedBlk.(*SignedBlock)
}
