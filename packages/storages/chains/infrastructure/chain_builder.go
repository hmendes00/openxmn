package infrastructure

import (
	"errors"

	stored_chained_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/chained/domain"
	stored_chains "github.com/XMNBlockchain/core/packages/storages/chains/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

type chainBuilder struct {
	met   stored_files.File
	ht    stored_files.File
	flBlk stored_chained_blocks.Block
	clBlk stored_chained_blocks.Block
}

func createChainBuilder() stored_chains.ChainBuilder {
	out := chainBuilder{
		met:   nil,
		ht:    nil,
		flBlk: nil,
		clBlk: nil,
	}

	return &out
}

// Create initializes the chain builder
func (build *chainBuilder) Create() stored_chains.ChainBuilder {
	build.met = nil
	build.ht = nil
	build.flBlk = nil
	build.clBlk = nil
	return build
}

// WithMetaData adds a metadata file to the chain builder
func (build *chainBuilder) WithMetaData(met stored_files.File) stored_chains.ChainBuilder {
	build.met = met
	return build
}

// WithHashTree adds an hashtree file to the chain builder
func (build *chainBuilder) WithHashTree(ht stored_files.File) stored_chains.ChainBuilder {
	build.ht = ht
	return build
}

// WithFloorBlock adds a floor block file to the chain builder
func (build *chainBuilder) WithFloorBlock(floorBlk stored_chained_blocks.Block) stored_chains.ChainBuilder {
	build.flBlk = floorBlk
	return build
}

// WithCeilBlock adds a ceil block file to the chain builder
func (build *chainBuilder) WithCeilBlock(ceilBlk stored_chained_blocks.Block) stored_chains.ChainBuilder {
	build.clBlk = ceilBlk
	return build
}

// Now builds a new Chain instance
func (build *chainBuilder) Now() (stored_chains.Chain, error) {
	if build.met == nil {
		return nil, errors.New("the metadata file is mandatory in order to build a chain instance")
	}

	if build.ht == nil {
		return nil, errors.New("the hashtree file is mandatory in order to build a chain instance")
	}

	if build.flBlk == nil {
		return nil, errors.New("the floor block file is mandatory in order to build a chain instance")
	}

	if build.clBlk == nil {
		return nil, errors.New("the ceil block file is mandatory in order to build a chain instance")
	}

	out := createChain(build.met, build.ht, build.flBlk, build.clBlk)
	return out, nil
}
