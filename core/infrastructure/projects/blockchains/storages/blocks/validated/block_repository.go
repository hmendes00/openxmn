package validated

import (
	"path/filepath"

	stored_block "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks"
	stored_validated_block "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks/validated"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/users"
)

// BlockRepository represents a concrete stored validated BlockRepository implementation
type BlockRepository struct {
	fileRepository             stored_files.FileRepository
	signedBlkRepository        stored_block.SignedBlockRepository
	sigsRepository             stored_users.SignaturesRepository
	validatedBlkBuilderFactory stored_validated_block.BlockBuilderFactory
}

// CreateBlockRepository creates a new BlockRepository instance
func CreateBlockRepository(fileRepository stored_files.FileRepository, signedBlkRepository stored_block.SignedBlockRepository, sigsRepository stored_users.SignaturesRepository, validatedBlkBuilderFactory stored_validated_block.BlockBuilderFactory) stored_validated_block.BlockRepository {
	out := BlockRepository{
		fileRepository:             fileRepository,
		signedBlkRepository:        signedBlkRepository,
		sigsRepository:             sigsRepository,
		validatedBlkBuilderFactory: validatedBlkBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a stored validated block instance
func (rep *BlockRepository) Retrieve(dirPath string) (stored_validated_block.Block, error) {
	metPath := filepath.Join(dirPath, "metadata.json")
	met, metErr := rep.fileRepository.Retrieve(metPath)
	if metErr != nil {
		return nil, metErr
	}

	blkPath := filepath.Join(dirPath, "signed_block")
	signedBlk, signedBlkErr := rep.signedBlkRepository.Retrieve(blkPath)
	if signedBlkErr != nil {
		return nil, signedBlkErr
	}

	sigsPath := filepath.Join(dirPath, "signatures")
	sigs, sigsErr := rep.sigsRepository.Retrieve(sigsPath)
	if sigsErr != nil {
		return nil, sigsErr
	}

	out, outErr := rep.validatedBlkBuilderFactory.Create().Create().WithMetaData(met).WithSignatures(sigs).WithBlock(signedBlk).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
