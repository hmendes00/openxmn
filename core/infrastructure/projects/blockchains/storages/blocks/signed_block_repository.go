package blocks

import (
	"path/filepath"

	stored_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/users"
)

// SignedBlockRepository represents a concrete stored signed block repository implementation
type SignedBlockRepository struct {
	fileRepository          stored_files.FileRepository
	sigRepository           stored_users.SignatureRepository
	blkRepository           stored_blocks.BlockRepository
	signedBlkBuilderFactory stored_blocks.SignedBlockBuilderFactory
}

// CreateSignedBlockRepository creates a new SignedBlockRepository instance
func CreateSignedBlockRepository(fileRepository stored_files.FileRepository, sigRepository stored_users.SignatureRepository, blkRepository stored_blocks.BlockRepository, signedBlkBuilderFactory stored_blocks.SignedBlockBuilderFactory) stored_blocks.SignedBlockRepository {
	out := SignedBlockRepository{
		fileRepository:          fileRepository,
		sigRepository:           sigRepository,
		blkRepository:           blkRepository,
		signedBlkBuilderFactory: signedBlkBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a signed block instance
func (rep *SignedBlockRepository) Retrieve(dirPath string) (stored_blocks.SignedBlock, error) {
	metPath := filepath.Join(dirPath, "metadata.json")
	met, metErr := rep.fileRepository.Retrieve(metPath)
	if metErr != nil {
		return nil, metErr
	}

	sigPath := filepath.Join(dirPath, "signature")
	sig, sigErr := rep.sigRepository.Retrieve(sigPath)
	if sigErr != nil {
		return nil, sigErr
	}

	blkPath := filepath.Join(dirPath, "block")
	blk, blkErr := rep.blkRepository.Retrieve(blkPath)
	if blkErr != nil {
		return nil, blkErr
	}

	out, outErr := rep.signedBlkBuilderFactory.Create().Create().WithMetaData(met).WithSignature(sig).WithBlock(blk).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
