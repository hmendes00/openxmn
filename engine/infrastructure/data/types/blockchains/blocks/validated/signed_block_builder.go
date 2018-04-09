package validated

import (
	"errors"
	"strconv"
	"time"

	validated "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks/validated"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
	uuid "github.com/satori/go.uuid"
)

type signedBlockBuilder struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.BuilderFactory
	id                     *uuid.UUID
	met                    metadata.MetaData
	blk                    validated.Block
	sig                    users.Signature
	crOn                   *time.Time
}

func createSignedBlockBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.BuilderFactory) validated.SignedBlockBuilder {
	out := signedBlockBuilder{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
		id:   nil,
		met:  nil,
		blk:  nil,
		sig:  nil,
		crOn: nil,
	}

	return &out
}

// Create initializes the SignedBlockBuilder instance
func (build *signedBlockBuilder) Create() validated.SignedBlockBuilder {
	build.id = nil
	build.met = nil
	build.blk = nil
	build.sig = nil
	build.crOn = nil
	return build
}

// WithID adds an ID to the SignedBlockBuilder instance
func (build *signedBlockBuilder) WithID(id *uuid.UUID) validated.SignedBlockBuilder {
	build.id = id
	return build
}

// WithMetaData adds metadata to the SignedBlockBuilder instance
func (build *signedBlockBuilder) WithMetaData(met metadata.MetaData) validated.SignedBlockBuilder {
	build.met = met
	return build
}

// WithBlock adds a block to the SignedBlockBuilder instance
func (build *signedBlockBuilder) WithBlock(blk validated.Block) validated.SignedBlockBuilder {
	build.blk = blk
	return build
}

// WithSignature adds a signature to the SignedBlockBuilder instance
func (build *signedBlockBuilder) WithSignature(sig users.Signature) validated.SignedBlockBuilder {
	build.sig = sig
	return build
}

// CreatedOn adds a creation time to the SignedBlockBuilder instance
func (build *signedBlockBuilder) CreatedOn(ts time.Time) validated.SignedBlockBuilder {
	build.crOn = &ts
	return build
}

// Now builds a new SignedBlock instance
func (build *signedBlockBuilder) Now() (validated.SignedBlock, error) {

	if build.blk == nil {
		return nil, errors.New("the block is mandatory in order to build a SignedBlock instance")
	}

	if build.sig == nil {
		return nil, errors.New("the signature is mandatory in order to build a SignedBlock instance")
	}

	if build.met == nil {
		if build.id == nil {
			return nil, errors.New("the ID is mandatory in order to build a validated block")
		}

		if build.crOn == nil {
			return nil, errors.New("the creation time is mandatory in order to build a validated block")
		}

		blocks := [][]byte{
			build.id.Bytes(),
			[]byte(strconv.Itoa(int(build.crOn.UnixNano()))),
			build.blk.GetMetaData().GetHashTree().GetHash().Get(),
			build.sig.GetMetaData().GetID().Bytes(),
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

	out := createSignedBlock(build.met.(*concrete_metadata.MetaData), build.blk.(*Block), build.sig.(*concrete_users.Signature))
	return out, nil
}
