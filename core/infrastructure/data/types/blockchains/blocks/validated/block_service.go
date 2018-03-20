package validated

import (
	"path/filepath"

	blocks "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/blocks"
	validated "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/blocks/validated"
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/metadata"
	stored_validated_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/blocks/validated"
	users "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/users"
)

// BlockService represents a concrete BlockService implementation
type BlockService struct {
	metaDataService                  metadata.MetaDataService
	signedBlkService                 blocks.SignedBlockService
	userSigService                   users.SignaturesService
	storedValidatedBlkBuilderFactory stored_validated_blocks.BlockBuilderFactory
}

// CreateBlockService creates a new BlockService instance
func CreateBlockService(
	metaDataService metadata.MetaDataService,
	signedBlkService blocks.SignedBlockService,
	userSigService users.SignaturesService,
	storedValidatedBlkBuilderFactory stored_validated_blocks.BlockBuilderFactory,
) validated.BlockService {
	out := BlockService{
		metaDataService:                  metaDataService,
		signedBlkService:                 signedBlkService,
		userSigService:                   userSigService,
		storedValidatedBlkBuilderFactory: storedValidatedBlkBuilderFactory,
	}
	return &out
}

// Save saves a Block
func (serv *BlockService) Save(dirPath string, validatedBlk validated.Block) (stored_validated_blocks.Block, error) {
	//save the metadata:
	met := validatedBlk.GetMetaData()
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//save the block:
	blk := validatedBlk.GetBlock()
	blkPath := filepath.Join(dirPath, "signed_block")
	storedBlk, storedBlkErr := serv.signedBlkService.Save(blkPath, blk)
	if storedBlkErr != nil {
		return nil, storedBlkErr
	}

	//save the user signatures:
	sigs := validatedBlk.GetSignatures()
	sigsPath := filepath.Join(dirPath, "signatures")
	storedSigs, storedSigsErr := serv.userSigService.Save(sigsPath, sigs)
	if storedSigsErr != nil {
		return nil, storedSigsErr
	}

	//build the stored validated block:
	storedValidatedBlk, storedValidatedBlkErr := serv.storedValidatedBlkBuilderFactory.Create().Create().WithMetaData(storedMet).WithBlock(storedBlk).WithSignatures(storedSigs).Now()
	if storedValidatedBlkErr != nil {
		return nil, storedValidatedBlkErr
	}

	return storedValidatedBlk, nil
}
