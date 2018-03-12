package blocks

import (
	"errors"

	stored_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	stored_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/transactions/signed/aggregated"
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

	out := createBlock(build.met, build.trs)
	return out, nil
}
