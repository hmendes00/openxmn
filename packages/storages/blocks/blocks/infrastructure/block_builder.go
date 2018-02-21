package infrastructure

import (
	"errors"

	stored_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/blocks/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_aggregated_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/aggregated/domain"
)

type blockBuilder struct {
	metadata stored_files.File
	ht       stored_files.File
	trs      []stored_aggregated_transactions.SignedTransactions
}

func createBlockBuilder() stored_blocks.BlockBuilder {
	out := blockBuilder{
		metadata: nil,
		ht:       nil,
		trs:      nil,
	}

	return &out
}

// Create initializes the BlockBuilder instance
func (build *blockBuilder) Create() stored_blocks.BlockBuilder {
	build.metadata = nil
	build.ht = nil
	build.trs = nil
	return build
}

// WithMetaData adds metadata to the BlockBuilder instance
func (build *blockBuilder) WithMetaData(met stored_files.File) stored_blocks.BlockBuilder {
	build.metadata = met
	return build
}

// WithHashTree adds an hashtree to the BlockBuilder instance
func (build *blockBuilder) WithHashTree(ht stored_files.File) stored_blocks.BlockBuilder {
	build.ht = ht
	return build
}

// WithTransactions adds stored transactions to the BlockBuilder instance
func (build *blockBuilder) WithTransactions(trs []stored_aggregated_transactions.SignedTransactions) stored_blocks.BlockBuilder {
	build.trs = trs
	return build
}

// Now builds a new Block instance
func (build *blockBuilder) Now() (stored_blocks.Block, error) {
	if build.metadata == nil {
		return nil, errors.New("the metadata is mandatory in order to build a Block instance")
	}

	if build.ht == nil {
		return nil, errors.New("the hashtree is mandatory in order to build a Block instance")
	}

	if build.trs == nil {
		return nil, errors.New("the transactions is mandatory in order to build a Block instance")
	}

	out := createBlock(build.metadata, build.ht, build.trs)
	return out, nil
}
