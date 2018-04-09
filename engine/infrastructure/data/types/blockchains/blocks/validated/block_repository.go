package validated

import (
	"path/filepath"

	blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks"
	validated "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks/validated"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// BlockRepository represents a concrete BlockRepository implementation
type BlockRepository struct {
	metaDataRepository         metadata.Repository
	signedBlkRepository        blocks.SignedBlockRepository
	userSigRepository          users.SignaturesRepository
	validatedBlkBuilderFactory validated.BlockBuilderFactory
}

// CreateBlockRepository creates a new BlockRepository instance
func CreateBlockRepository(
	metaDataRepository metadata.Repository,
	signedBlkRepository blocks.SignedBlockRepository,
	userSigRepository users.SignaturesRepository,
	validatedBlkBuilderFactory validated.BlockBuilderFactory,
) validated.BlockRepository {
	out := BlockRepository{
		metaDataRepository:         metaDataRepository,
		signedBlkRepository:        signedBlkRepository,
		userSigRepository:          userSigRepository,
		validatedBlkBuilderFactory: validatedBlkBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a validated block
func (rep *BlockRepository) Retrieve(dirPath string) (validated.Block, error) {
	//retrieve the metadata:
	met, metErr := rep.metaDataRepository.Retrieve(dirPath)
	if metErr != nil {
		return nil, metErr
	}

	//retrieve the signed block:
	blkPath := filepath.Join(dirPath, "signed_block")
	signedBlk, signedBlkErr := rep.signedBlkRepository.Retrieve(blkPath)
	if signedBlkErr != nil {
		return nil, signedBlkErr
	}

	//retrieve the user signatures:
	sigsPath := filepath.Join(dirPath, "signatures")
	userSigs, userSigsErr := rep.userSigRepository.Retrieve(sigsPath)
	if userSigsErr != nil {
		return nil, userSigsErr
	}

	//build the block:
	id := met.GetID()
	ts := met.CreatedOn()
	validatedBlk, validatedBlkErr := rep.validatedBlkBuilderFactory.Create().Create().WithID(id).CreatedOn(ts).WithBlock(signedBlk).WithSignatures(userSigs).Now()
	if validatedBlkErr != nil {
		return nil, validatedBlkErr
	}

	return validatedBlk, nil
}
