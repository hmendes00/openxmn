package chained

import (
	"errors"
	"strconv"
	"time"

	validated "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/blocks/validated"
	chained "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/blocks/validated/chained"
	hashtrees "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/hashtrees"
	concrete_validated "github.com/XMNBlockchain/exmachina-network/engine/infrastructure/data/types/blockchains/blocks/validated"
	uuid "github.com/satori/go.uuid"
)

type blockBuilder struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory chained.MetaDataBuilderFactory
	id                     *uuid.UUID
	met                    chained.MetaData
	blk                    validated.Block
	prevID                 *uuid.UUID
	crOn                   *time.Time
}

func createBlockBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory chained.MetaDataBuilderFactory) chained.BlockBuilder {
	out := blockBuilder{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
		id:     nil,
		met:    nil,
		blk:    nil,
		prevID: nil,
		crOn:   nil,
	}

	return &out
}

// Create initializes the BlockBuilder instance
func (build *blockBuilder) Create() chained.BlockBuilder {
	build.id = nil
	build.met = nil
	build.blk = nil
	build.prevID = nil
	build.crOn = nil
	return build
}

// WithID adds an ID to the BlockBuilder
func (build *blockBuilder) WithID(id *uuid.UUID) chained.BlockBuilder {
	build.id = id
	return build
}

// WithMetaData adds a MetaData instance to the BlockBuilder
func (build *blockBuilder) WithMetaData(met chained.MetaData) chained.BlockBuilder {
	build.met = met
	return build
}

// WithBlock adds a Block instance to the BlockBuilder
func (build *blockBuilder) WithBlock(blk validated.Block) chained.BlockBuilder {
	build.blk = blk
	return build
}

// WithPreviousID adds a previous ID to the BlockBuilder
func (build *blockBuilder) WithPreviousID(prevID *uuid.UUID) chained.BlockBuilder {
	build.prevID = prevID
	return build
}

// CreatedOn adds a creation time to the BlockBuilder
func (build *blockBuilder) CreatedOn(crOn time.Time) chained.BlockBuilder {
	build.crOn = &crOn
	return build
}

// Now builds a Block instance
func (build *blockBuilder) Now() (chained.Block, error) {

	if build.blk == nil {
		return nil, errors.New("the validated block is mandatory in order to build a Block instance")
	}

	if build.met == nil {

		if build.id == nil {
			return nil, errors.New("the ID is mandatory in order to build a Block instance")
		}

		if build.prevID == nil {
			return nil, errors.New("the previous ID is mandatory in order to build a Block instance")
		}

		if build.crOn == nil {
			return nil, errors.New("the previous ID is mandatory in order to build a Block instance")
		}

		blocks := [][]byte{
			build.id.Bytes(),
			[]byte(strconv.Itoa(int(build.crOn.UnixNano()))),
			build.prevID.Bytes(),
			build.blk.GetMetaData().GetHashTree().GetHash().Get(),
		}

		ht, htErr := build.htBuilderFactory.Create().Create().WithBlocks(blocks).Now()
		if htErr != nil {
			return nil, htErr
		}

		met, metErr := build.metaDataBuilderFactory.Create().Create().WithID(build.id).WithPreviousID(build.prevID).WithHashTree(ht).CreatedOn(*build.crOn).Now()
		if metErr != nil {
			return nil, metErr
		}

		build.met = met

	}

	if build.met == nil {
		return nil, errors.New("the MetaData is mandatory in order to build a Block instance")
	}

	out := createBlock(build.met.(*MetaData), build.blk.(*concrete_validated.Block))
	return out, nil
}
