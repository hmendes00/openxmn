package infrastructure

import (
	"errors"
	"path/filepath"

	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
)

// SignedBlockRepository represents a concrete SignedBlockRepository implementation
type SignedBlockRepository struct {
	objRepository           objects.ObjectRepository
	blkRepository           blocks.BlockRepository
	signedBlkBuilderFactory blocks.SignedBlockBuilderFactory
}

// CreateSignedBlockRepository creates a new SignedBlockRepository instance
func CreateSignedBlockRepository(
	objRepository objects.ObjectRepository,
	blkRepository blocks.BlockRepository,
	signedBlkBuilderFactory blocks.SignedBlockBuilderFactory,
) blocks.SignedBlockRepository {
	out := SignedBlockRepository{
		objRepository:           objRepository,
		blkRepository:           blkRepository,
		signedBlkBuilderFactory: signedBlkBuilderFactory,
	}
	return &out
}

// Retrieve retrieves a SignedBlock instance
func (rep *SignedBlockRepository) Retrieve(dirPath string) (blocks.SignedBlock, error) {
	//retrieve the object:
	obj, objErr := rep.objRepository.Retrieve(dirPath)
	if objErr != nil {
		return nil, objErr
	}

	//retrieve the block:
	blkPath := filepath.Join(dirPath, "block")
	blk, blkErr := rep.blkRepository.Retrieve(blkPath)
	if blkErr != nil {
		return nil, blkErr
	}

	//there must be a signature inside the metadata:
	met := obj.GetMetaData()
	if !met.HasSignature() {
		return nil, errors.New("the signature is mandatory, inside the metadata of the object, in order to retrieve a SignedBlock instance")
	}

	//build the signed block:
	id := met.GetID()
	sig := met.GetSignature()
	ts := met.CreatedOn()
	signedBlk, signedBlkErr := rep.signedBlkBuilderFactory.Create().Create().WithID(id).WithSignature(sig).WithBlock(blk).CreatedOn(ts).Now()
	if signedBlkErr != nil {
		return nil, signedBlkErr
	}

	return signedBlk, nil
}
