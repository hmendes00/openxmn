package infrastructure

import (
	"path/filepath"

	blocks "github.com/XMNBlockchain/core/packages/blockchains/blocks/blocks/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
	stored_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/blocks/domain"
)

// BlockService represents a concrete BlockService implementation
type BlockService struct {
	metaDataService   metadata.MetaDataService
	signedTrsService  aggregated.SignedTransactionsService
	blkBuilderFactory stored_blocks.BlockBuilderFactory
}

// CreateBlockService creates a new BlockService instance
func CreateBlockService(metaDataService metadata.MetaDataService, signedTrsService aggregated.SignedTransactionsService, blkBuilderFactory stored_blocks.BlockBuilderFactory) blocks.BlockService {
	out := BlockService{
		metaDataService:   metaDataService,
		signedTrsService:  signedTrsService,
		blkBuilderFactory: blkBuilderFactory,
	}

	return &out
}

// Save saves an Block instance
func (serv *BlockService) Save(dirPath string, trs blocks.Block) (stored_blocks.Block, error) {
	//save the metaData
	met := trs.GetMetaData()
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//save the transactions:
	signedTrs := trs.GetTransactions()
	trsPath := filepath.Join(dirPath, "signed_transactions")
	storedTrs, storedTrsErr := serv.signedTrsService.SaveAll(trsPath, signedTrs)
	if storedTrsErr != nil {
		return nil, storedTrsErr
	}

	//build the block:
	out, outErr := serv.blkBuilderFactory.Create().Create().WithMetaData(storedMet).WithTransactions(storedTrs).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
