package validated

import (
	"path/filepath"

	validated "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks/validated"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// SignedBlockRepository represents a SignedBlockRepository implementation
type SignedBlockRepository struct {
	metaDataRepository      metadata.Repository
	blkRepository           validated.BlockRepository
	sigRepository           users.SignatureRepository
	signedBlkBuilderFactory validated.SignedBlockBuilderFactory
}

// CreateSignedBlockRepository creates a new SignedBlockRepository instance
func CreateSignedBlockRepository(metaDataRepository metadata.Repository, blkRepository validated.BlockRepository, sigRepository users.SignatureRepository, signedBlkBuilderFactory validated.SignedBlockBuilderFactory) validated.SignedBlockRepository {
	out := SignedBlockRepository{
		metaDataRepository:      metaDataRepository,
		blkRepository:           blkRepository,
		sigRepository:           sigRepository,
		signedBlkBuilderFactory: signedBlkBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a SignedBlock instance
func (rep *SignedBlockRepository) Retrieve(dirPath string) (validated.SignedBlock, error) {
	//retrieve the metadata:
	met, metErr := rep.metaDataRepository.Retrieve(dirPath)
	if metErr != nil {
		return nil, metErr
	}

	//retrieve the block:
	blkPath := filepath.Join(dirPath, "validated_block")
	blk, blkErr := rep.blkRepository.Retrieve(blkPath)
	if blkErr != nil {
		return nil, blkErr
	}

	//retrieve the signature:
	sigPath := filepath.Join(dirPath, "signature")
	sig, sigErr := rep.sigRepository.Retrieve(sigPath)
	if sigErr != nil {
		return nil, sigErr
	}

	out, outErr := rep.signedBlkBuilderFactory.Create().Create().WithMetaData(met).WithSignature(sig).WithBlock(blk).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
