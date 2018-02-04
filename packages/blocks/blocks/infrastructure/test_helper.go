package infrastructure

import (
	"math/rand"
	"testing"

	concrete_hashtree "github.com/XMNBlockchain/core/packages/hashtrees/infrastructure"
	concrete_aggregated "github.com/XMNBlockchain/core/packages/transactions/aggregated/infrastructure"
)

// CreateBlockForTests creates a Block for tests
func CreateBlockForTests(t *testing.T) *Block {
	//variables:
	trs := []*concrete_aggregated.SignedTransactions{
		concrete_aggregated.CreateSignedTransactionsForTests(t),
		concrete_aggregated.CreateSignedTransactionsForTests(t),
		concrete_aggregated.CreateSignedTransactionsForTests(t),
	}

	ht := concrete_hashtree.CreateHashTreeForTests(t)
	neededKarma := rand.Int()%500 + 100

	blk := createBlock(ht, trs, neededKarma)
	return blk.(*Block)
}
