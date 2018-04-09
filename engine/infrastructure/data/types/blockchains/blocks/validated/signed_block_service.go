package validated

import (
	"path/filepath"

	stored_validated_blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/blocks/validated"
	validated "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks/validated"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// SignedBlockService represents a concrete SignedBlockService implementation
type SignedBlockService struct {
	metaDataService                        metadata.Service
	blkService                             validated.BlockService
	signatureService                       users.SignatureService
	storedSignedValidatedBlkBuilderFactory stored_validated_blocks.SignedBlockBuilderFactory
}

// CreateSignedBlockService creates a new SignedBlockService instance
func CreateSignedBlockService(metaDataService metadata.Service, blkService validated.BlockService, signatureService users.SignatureService, storedSignedValidatedBlkBuilderFactory stored_validated_blocks.SignedBlockBuilderFactory) validated.SignedBlockService {
	out := SignedBlockService{
		metaDataService:                        metaDataService,
		blkService:                             blkService,
		signatureService:                       signatureService,
		storedSignedValidatedBlkBuilderFactory: storedSignedValidatedBlkBuilderFactory,
	}

	return &out
}

// Save saves a SignedBlock instance
func (serv *SignedBlockService) Save(dirPath string, signedBlk validated.SignedBlock) (stored_validated_blocks.SignedBlock, error) {
	//save the metadata:
	met := signedBlk.GetMetaData()
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//save the block:
	blk := signedBlk.GetBlock()
	blkPath := filepath.Join(dirPath, "validated_block")
	storedBlk, storedBlkErr := serv.blkService.Save(blkPath, blk)
	if storedBlkErr != nil {
		return nil, storedBlkErr
	}

	//save the signature:
	sig := signedBlk.GetSignature()
	sigPath := filepath.Join(dirPath, "signature")
	storedSig, storedSigErr := serv.signatureService.Save(sigPath, sig)
	if storedSigErr != nil {
		return nil, storedSigErr
	}

	//build the output:
	out, outErr := serv.storedSignedValidatedBlkBuilderFactory.Create().Create().WithMetaData(storedMet).WithBlock(storedBlk).WithSignature(storedSig).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
