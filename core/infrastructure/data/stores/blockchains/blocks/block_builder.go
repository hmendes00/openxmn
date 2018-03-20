package blocks

import (
	"errors"

	stored_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/blocks"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/files"
	stored_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/transactions/signed/aggregated"
	conrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/stores/files"
	concrete_stored_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/stores/blockchains/transactions/signed/aggregated"
)

type blockBuilder struct {
	met stored_files.File
	trs []stored_aggregated_transactions.SignedTransactions
}

func createBlockBuilder() stored_blocks.BlockBuilder {
	out := blockBuilder{
		met: nil,
		trs: nil,
	}

	return &out
}

// Create initializes the BlockBuilder
func (build *blockBuilder) Create() stored_blocks.BlockBuilder {
	build.met = nil
	build.trs = nil
	return build
}

// WithMetaData adds MetaData to the BlockBuilder instance
func (build *blockBuilder) WithMetaData(met stored_files.File) stored_blocks.BlockBuilder {
	build.met = met
	return build
}

// WithTransactions adds SignedTransactions to the BlockBuilder instance
func (build *blockBuilder) WithTransactions(trs []stored_aggregated_transactions.SignedTransactions) stored_blocks.BlockBuilder {
	build.trs = trs
	return build
}

// Now builds a new Block instance
func (build *blockBuilder) Now() (stored_blocks.Block, error) {
	if build.met == nil {
		return nil, errors.New("the MetaData is mandatory in order to build a Block instance")
	}

	if build.trs == nil {
		return nil, errors.New("the SignedTransactions are mandatory in order to build a Block instance")
	}

	trs := []*concrete_stored_aggregated_transactions.SignedTransactions{}
	for _, oneTrs := range build.trs {
		trs = append(trs, oneTrs.(*concrete_stored_aggregated_transactions.SignedTransactions))
	}

	out := createBlock(build.met.(*conrete_stored_files.File), trs)
	return out, nil
}
