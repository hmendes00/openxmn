package infrastructure

import (
	stored_chained_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/chained/domain"
	stored_chains "github.com/XMNBlockchain/core/packages/storages/chains/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

type chain struct {
	met   stored_files.File
	ht    stored_files.File
	flBlk stored_chained_blocks.Block
	clBlk stored_chained_blocks.Block
}

func createChain(met stored_files.File, ht stored_files.File, flBlk stored_chained_blocks.Block, clBlk stored_chained_blocks.Block) stored_chains.Chain {
	out := chain{
		met:   met,
		ht:    ht,
		flBlk: flBlk,
		clBlk: clBlk,
	}

	return &out
}

// GetMetaData returns the metadata file
func (ch *chain) GetMetaData() stored_files.File {
	return ch.met
}

// GetHashTree returns the hashtree file
func (ch *chain) GetHashTree() stored_files.File {
	return ch.ht
}

// GetFloorBlock returns the stored floor block
func (ch *chain) GetFloorBlock() stored_chained_blocks.Block {
	return ch.flBlk
}

// GetCeilBlock returns the stored ceil block
func (ch *chain) GetCeilBlock() stored_chained_blocks.Block {
	return ch.clBlk
}
