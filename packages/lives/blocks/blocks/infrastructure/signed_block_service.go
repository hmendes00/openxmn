package infrastructure

import (
	"path/filepath"

	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// SignedBlockService represents a concrete SignedBlockService implementation
type SignedBlockService struct {
	storedTreeBuilderFactory stored_objects.TreeBuilderFactory
	metaDataBuilderFactory   objects.MetaDataBuilderFactory
	objBuilderFactory        objects.ObjectBuilderFactory
	objService               objects.ObjectService
	blkService               blocks.BlockService
}

// CreateSignedBlockService creates a new SignedBlockService instance
func CreateSignedBlockService(
	storedTreeBuilderFactory stored_objects.TreeBuilderFactory,
	metaDataBuilderFactory objects.MetaDataBuilderFactory,
	objBuilderFactory objects.ObjectBuilderFactory,
	objService objects.ObjectService,
	blkService blocks.BlockService,
) blocks.SignedBlockService {
	out := SignedBlockService{
		storedTreeBuilderFactory: storedTreeBuilderFactory,
		metaDataBuilderFactory:   metaDataBuilderFactory,
		objBuilderFactory:        objBuilderFactory,
		objService:               objService,
		blkService:               blkService,
	}
	return &out
}

// Save saves a block instance
func (serv *SignedBlockService) Save(dirPath string, signedBlk blocks.SignedBlock) (stored_objects.Tree, error) {
	//build the metadata:
	id := signedBlk.GetID()
	sig := signedBlk.GetSignature()
	ts := signedBlk.CreatedOn()
	met, metErr := serv.metaDataBuilderFactory.Create().Create().WithID(id).WithSignature(sig).CreatedOn(ts).Now()
	if metErr != nil {
		return nil, metErr
	}

	//build the object:
	obj, objErr := serv.objBuilderFactory.Create().Create().WithMetaData(met).Now()
	if objErr != nil {
		return nil, objErr
	}

	//save the object:
	storedObj, storedObjErr := serv.objService.Save(dirPath, obj)
	if storedObjErr != nil {
		return nil, storedObjErr
	}

	//save the block:
	blk := signedBlk.GetBlock()
	blkPath := filepath.Join(dirPath, "block")
	storedBlk, storedBlkErr := serv.blkService.Save(blkPath, blk)
	if storedBlkErr != nil {
		return nil, storedBlkErr
	}

	//build the stored tree:
	storedTree, storedTreeErr := serv.storedTreeBuilderFactory.Create().Create().WithName("signed_block").WithObject(storedObj).WithSubTree(storedBlk).Now()
	if storedTreeErr != nil {
		return nil, storedTreeErr
	}

	return storedTree, nil
}
