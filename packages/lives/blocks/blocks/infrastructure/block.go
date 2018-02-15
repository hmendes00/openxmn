package infrastructure

import (
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/domain"
	concrete_hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/infrastructure"
	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/domain"
	concrete_aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// Block represents a concrete Block implementation
type Block struct {
	ID          *uuid.UUID                                `json:"id"`
	HT          *concrete_hashtrees.HashTree              `json:"hashtree"`
	Trs         []*concrete_aggregated.SignedTransactions `json:"transactions"`
	NeededKarma int                                       `json:"needed_karma"`
	CrOn        time.Time                                 `json:"created_on"`
}

func createBlock(id *uuid.UUID, ht *concrete_hashtrees.HashTree, trs []*concrete_aggregated.SignedTransactions, neededKarma int, crOn time.Time) blocks.Block {
	out := Block{
		ID:          id,
		HT:          ht,
		Trs:         trs,
		NeededKarma: neededKarma,
		CrOn:        crOn,
	}

	return &out
}

// GetID returns the ID of the block
func (blk *Block) GetID() *uuid.UUID {
	return blk.ID
}

// GetHashTree returns the HashTree
func (blk *Block) GetHashTree() hashtrees.HashTree {
	return blk.HT
}

// GetNeededKarma returns the needed karma to validate this block
func (blk *Block) GetNeededKarma() int {
	return blk.NeededKarma
}

// GetTransactions returns the Signed Transactions
func (blk *Block) GetTransactions() []aggregated.SignedTransactions {
	out := []aggregated.SignedTransactions{}
	for _, oneSignedTrs := range blk.Trs {
		out = append(out, oneSignedTrs)
	}

	return out
}

// CreatedOn returns the creation time of the block
func (blk *Block) CreatedOn() time.Time {
	return blk.CrOn
}
