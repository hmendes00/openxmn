package blocks

import (
	"path/filepath"

	stored_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/blocks"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/files"
	stored_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/transactions/signed/aggregated"
)

// BlockRepository represents a concrete stored BlockRepository implementation
type BlockRepository struct {
	fileRepository      stored_files.FileRepository
	aggrTransRepository stored_aggregated_transactions.SignedTransactionsRepository
	blkBuilderFactory   stored_blocks.BlockBuilderFactory
}

// CreateBlockRepository creates a new BlockRepository instance
func CreateBlockRepository(fileRepository stored_files.FileRepository, aggrTransRepository stored_aggregated_transactions.SignedTransactionsRepository, blkBuilderFactory stored_blocks.BlockBuilderFactory) stored_blocks.BlockRepository {
	out := BlockRepository{
		fileRepository:      fileRepository,
		aggrTransRepository: aggrTransRepository,
		blkBuilderFactory:   blkBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a stored block instance
func (rep *BlockRepository) Retrieve(dirPath string) (stored_blocks.Block, error) {
	metPath := filepath.Join(dirPath, "metadata.json")
	met, metErr := rep.fileRepository.Retrieve(metPath)
	if metErr != nil {
		return nil, metErr
	}

	trsPath := filepath.Join(dirPath, "signed_transactions")
	aggrTrans, aggrTransErr := rep.aggrTransRepository.RetrieveAll(trsPath)
	if aggrTransErr != nil {
		return nil, aggrTransErr
	}

	out, outErr := rep.blkBuilderFactory.Create().Create().WithMetaData(met).WithTransactions(aggrTrans).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
