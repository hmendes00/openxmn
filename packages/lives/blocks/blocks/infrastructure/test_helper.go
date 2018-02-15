package infrastructure

import (
	"math/rand"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/infrastructure"
	concrete_aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/infrastructure"
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

	ht := concrete_hashtrees.CreateHashTreeForTests(t)

	neededKarma := rand.Int()%500 + 100

	blk := createBlock(&id, ht, trs, neededKarma, crOn)
	return blk.(*Block)
}
