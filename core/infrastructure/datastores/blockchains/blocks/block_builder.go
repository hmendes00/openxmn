package blocks

import (
	"errors"
	"strconv"
	"time"

	blocks "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/blocks"
	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/hashtrees"
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/metadata"
	aggregated "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/transactions/signed/aggregated"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/metadata"
	concrete_aggregated "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/transactions/signed/aggregated"
	uuid "github.com/satori/go.uuid"
)

type blockBuilder struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.MetaDataBuilderFactory
	id                     *uuid.UUID
	met                    metadata.MetaData
	trs                    []aggregated.SignedTransactions
	crOn                   *time.Time
}

func createBlockBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.MetaDataBuilderFactory) blocks.BlockBuilder {
	out := blockBuilder{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
		id:   nil,
		met:  nil,
		trs:  nil,
		crOn: nil,
	}

	return &out
}

// Create initializes the BlockBuilder instance
func (build *blockBuilder) Create() blocks.BlockBuilder {
	build.id = nil
	build.met = nil
	build.trs = nil
	build.crOn = nil
	return build
}

// WithID adds an ID to the BlockBuilder instance
func (build *blockBuilder) WithID(id *uuid.UUID) blocks.BlockBuilder {
	build.id = id
	return build
}

// WithMetaData adds MetaData to the BlockBuilder instance
func (build *blockBuilder) WithMetaData(met metadata.MetaData) blocks.BlockBuilder {
	build.met = met
	return build
}

// WithTransactions adds []SignedTransactions to the BlockBuilder instance
func (build *blockBuilder) WithTransactions(trs []aggregated.SignedTransactions) blocks.BlockBuilder {
	build.trs = trs
	return build
}

// CreatedOn adds creation time to the BlockBuilder instance
func (build *blockBuilder) CreatedOn(crOn time.Time) blocks.BlockBuilder {
	build.crOn = &crOn
	return build
}

// Now builds a new Block instance
func (build *blockBuilder) Now() (blocks.Block, error) {
	if build.trs == nil {
		return nil, errors.New("the MetaData is mandatory in order to build an Block instance ")
	}

	if len(build.trs) <= 0 {
		return nil, errors.New("the amount of SignedTransactions must be greater than 0 in order to build an Block instance")
	}

	if build.met == nil {
		if build.id == nil {
			return nil, errors.New("the ID is mandatory in order to build an Block instance")
		}

		if build.crOn == nil {
			return nil, errors.New("the creation time is mandatory in order to build an Block instance")
		}

		blocks := [][]byte{
			build.id.Bytes(),
			[]byte(strconv.Itoa(int(build.crOn.UnixNano()))),
		}

		for _, oneTrs := range build.trs {
			blocks = append(blocks, oneTrs.GetMetaData().GetHashTree().GetHash().Get())
		}

		ht, htErr := build.htBuilderFactory.Create().Create().WithBlocks(blocks).Now()
		if htErr != nil {
			return nil, htErr
		}

		met, metErr := build.metaDataBuilderFactory.Create().Create().WithID(build.id).WithHashTree(ht).CreatedOn(*build.crOn).Now()
		if metErr != nil {
			return nil, metErr
		}

		build.met = met
	}

	if build.met == nil {
		return nil, errors.New("the MetaData is mandatory in order to build an Block instance")
	}

	trs := []*concrete_aggregated.SignedTransactions{}
	for _, oneTrs := range build.trs {
		trs = append(trs, oneTrs.(*concrete_aggregated.SignedTransactions))
	}

	out := createBlock(build.met.(*concrete_metadata.MetaData), trs)
	return out, nil
}
