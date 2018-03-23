package validated

import (
	"errors"
	"strconv"
	"time"

	blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks"
	validated "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks/validated"
	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	concrete_blocks "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/blocks"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
	uuid "github.com/satori/go.uuid"
)

type blockBuilder struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.MetaDataBuilderFactory
	id                     *uuid.UUID
	met                    metadata.MetaData
	blk                    blocks.SignedBlock
	ls                     users.Signatures
	ts                     *time.Time
}

func createBlockBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.MetaDataBuilderFactory) validated.BlockBuilder {
	out := blockBuilder{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
		id:  nil,
		met: nil,
		blk: nil,
		ls:  nil,
		ts:  nil,
	}

	return &out
}

// Create initializes the BlockBuilder instance
func (build *blockBuilder) Create() validated.BlockBuilder {
	build.id = nil
	build.met = nil
	build.blk = nil
	build.ls = nil
	build.ts = nil
	return build
}

// WithID adds an ID to the BlockBuilder instance
func (build *blockBuilder) WithID(id *uuid.UUID) validated.BlockBuilder {
	build.id = id
	return build
}

// WithMetaData adds MetaData to the BlockBuilder instance
func (build *blockBuilder) WithMetaData(met metadata.MetaData) validated.BlockBuilder {
	build.met = met
	return build
}

// WithBlock adds a Block to the BlockBuilder instance
func (build *blockBuilder) WithBlock(blk blocks.SignedBlock) validated.BlockBuilder {
	build.blk = blk
	return build
}

// WithSignatures adds a leader signatures to the BlockBuilder instance
func (build *blockBuilder) WithSignatures(sigs users.Signatures) validated.BlockBuilder {
	build.ls = sigs
	return build
}

// CreatedOn adds a creation time to the BlockBuilder instance
func (build *blockBuilder) CreatedOn(ts time.Time) validated.BlockBuilder {
	build.ts = &ts
	return build
}

// Now builds a new Block instance
func (build *blockBuilder) Now() (validated.Block, error) {

	if build.blk == nil {
		return nil, errors.New("the block is mandatory in order to build a validated block")
	}

	if build.ls == nil {
		return nil, errors.New("the leader signatures are mandatory in order to build a block instance")
	}

	if build.met == nil {
		if build.id == nil {
			return nil, errors.New("the ID is mandatory in order to build a validated block")
		}

		if build.ts == nil {
			return nil, errors.New("the creation time is mandatory in order to build a validated block")
		}

		blocks := [][]byte{
			build.id.Bytes(),
			[]byte(strconv.Itoa(int(build.ts.UnixNano()))),
			build.blk.GetMetaData().GetHashTree().GetHash().Get(),
			build.ls.GetMetaData().GetHashTree().GetHash().Get(),
		}

		ht, htErr := build.htBuilderFactory.Create().Create().WithBlocks(blocks).Now()
		if htErr != nil {
			return nil, htErr
		}

		met, metErr := build.metaDataBuilderFactory.Create().Create().WithID(build.id).WithHashTree(ht).CreatedOn(*build.ts).Now()
		if metErr != nil {
			return nil, metErr
		}

		build.met = met
	}

	if build.met == nil {
		return nil, errors.New("the MetaData is mandatory in order to build a Block instance")
	}

	out := createBlock(build.met.(*concrete_metadata.MetaData), build.blk.(*concrete_blocks.SignedBlock), build.ls.(*concrete_users.Signatures))
	return out, nil
}
