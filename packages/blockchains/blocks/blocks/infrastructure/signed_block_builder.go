package infrastructure

import (
	"errors"
	"strconv"
	"time"

	blocks "github.com/XMNBlockchain/core/packages/blockchains/blocks/blocks/domain"
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/blockchains/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

type signedBlockBuilder struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.MetaDataBuilderFactory
	id                     *uuid.UUID
	met                    metadata.MetaData
	blk                    blocks.Block
	sig                    users.Signature
	createdOn              *time.Time
}

func createSignedBlockBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.MetaDataBuilderFactory) blocks.SignedBlockBuilder {
	out := signedBlockBuilder{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
		id:        nil,
		met:       nil,
		blk:       nil,
		sig:       nil,
		createdOn: nil,
	}

	return &out
}

// Create initializes the SignedBlockBuilder instance
func (build *signedBlockBuilder) Create() blocks.SignedBlockBuilder {
	build.id = nil
	build.met = nil
	build.blk = nil
	build.sig = nil
	build.createdOn = nil
	return build
}

// WithID adds the ID to the SignedBlockBuilder instance
func (build *signedBlockBuilder) WithID(id *uuid.UUID) blocks.SignedBlockBuilder {
	build.id = id
	return build
}

// WithMetaData adds the MetaData to the SignedBlockBuilder instance
func (build *signedBlockBuilder) WithMetaData(met metadata.MetaData) blocks.SignedBlockBuilder {
	build.met = met
	return build
}

// WithBlock adds a Block to the SignedBlockBuilder instance
func (build *signedBlockBuilder) WithBlock(blk blocks.Block) blocks.SignedBlockBuilder {
	build.blk = blk
	return build
}

// WithSignature adds a user Signature to the SignedBlockBuilder instance
func (build *signedBlockBuilder) WithSignature(sig users.Signature) blocks.SignedBlockBuilder {
	build.sig = sig
	return build
}

// CreatedOn adds the creation time to the SignedBlockBuilder instance
func (build *signedBlockBuilder) CreatedOn(createdOn time.Time) blocks.SignedBlockBuilder {
	build.createdOn = &createdOn
	return build
}

// Now builds a SignedBlock instance
func (build *signedBlockBuilder) Now() (blocks.SignedBlock, error) {

	if build.blk == nil {
		return nil, errors.New("the block is mandatory in order to build a SignedBlock instance")
	}

	if build.sig == nil {
		return nil, errors.New("the user signature is mandatory in order to build a SignedBlock instance")
	}

	if build.met == nil {
		if build.id == nil {
			return nil, errors.New("the ID is mandatory in order to build a SignedBlock instance")
		}

		if build.createdOn == nil {
			return nil, errors.New("the creation time is mandatory in order to build a SignedBlock instance")
		}

		blocks := [][]byte{
			build.id.Bytes(),
			[]byte(strconv.Itoa(int(build.createdOn.UnixNano()))),
			build.blk.GetMetaData().GetHashTree().GetHash().Get(),
			build.sig.GetMetaData().GetHashTree().GetHash().Get(),
		}

		ht, htErr := build.htBuilderFactory.Create().Create().WithBlocks(blocks).Now()
		if htErr != nil {
			return nil, htErr
		}

		met, metErr := build.metaDataBuilderFactory.Create().Create().WithID(build.id).WithHashTree(ht).CreatedOn(*build.createdOn).Now()
		if metErr != nil {
			return nil, metErr
		}

		build.met = met
	}

	if build.met == nil {
		return nil, errors.New("the MetaData is mandatory in order to build a Block instance")
	}

	out := createSignedBlock(build.met.(*concrete_metadata.MetaData), build.blk.(*Block), build.sig.(*concrete_users.Signature))
	return out, nil
}
