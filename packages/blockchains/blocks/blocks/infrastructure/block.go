package infrastructure

import (
	"time"

	blocks "github.com/XMNBlockchain/core/packages/blockchains/blocks/blocks/domain"
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
	concrete_aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// Block represents a concrete Block implementation
type Block struct {
	ID   *uuid.UUID                                `json:"id"`
	HT   *concrete_hashtrees.HashTree              `json:"hashtree"`
	Trs  []*concrete_aggregated.SignedTransactions `json:"transactions"`
	CrOn time.Time                                 `json:"created_on"`
}

func createBlock(id *uuid.UUID, ht *concrete_hashtrees.HashTree, trs []*concrete_aggregated.SignedTransactions, crOn time.Time) blocks.Block {
	out := Block{
		ID:   id,
		HT:   ht,
		Trs:  trs,
		CrOn: crOn,
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
