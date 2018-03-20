package blocks

import (
	"path/filepath"

	blocks "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/blocks"
	metadata "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/metadata"
	users "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/users"
)

// SignedBlockRepository represents a concrete SignedBlockRepository implementation
type SignedBlockRepository struct {
	metaDataRepository      metadata.MetaDataRepository
	userSigRepository       users.SignatureRepository
	blkRepository           blocks.BlockRepository
	signedBlkBuilderFactory blocks.SignedBlockBuilderFactory
}

// CreateSignedBlockRepository creates a new SignedBlockRepository instance
func CreateSignedBlockRepository(
	metaDataRepository metadata.MetaDataRepository,
	userSigRepository users.SignatureRepository,
	blkRepository blocks.BlockRepository,
	signedBlkBuilderFactory blocks.SignedBlockBuilderFactory,
) blocks.SignedBlockRepository {
	out := SignedBlockRepository{
		metaDataRepository:      metaDataRepository,
		userSigRepository:       userSigRepository,
		blkRepository:           blkRepository,
		signedBlkBuilderFactory: signedBlkBuilderFactory,
	}
	return &out
}

// Retrieve retrieves a SignedBlock instance
func (rep *SignedBlockRepository) Retrieve(dirPath string) (blocks.SignedBlock, error) {
	//retrieve the metadata:
	met, metErr := rep.metaDataRepository.Retrieve(dirPath)
	if metErr != nil {
		return nil, metErr
	}

	//retrieve the signature:
	sigPath := filepath.Join(dirPath, "signature")
	sig, sigErr := rep.userSigRepository.Retrieve(sigPath)
	if sigErr != nil {
		return nil, sigErr
	}

	//retrieve the block:
	blkPath := filepath.Join(dirPath, "block")
	blk, blkErr := rep.blkRepository.Retrieve(blkPath)
	if blkErr != nil {
		return nil, blkErr
	}

	//build the signed block:
	signedBlk, signedBlkErr := rep.signedBlkBuilderFactory.Create().Create().WithMetaData(met).WithSignature(sig).WithBlock(blk).Now()
	if signedBlkErr != nil {
		return nil, signedBlkErr
	}

	return signedBlk, nil
}
