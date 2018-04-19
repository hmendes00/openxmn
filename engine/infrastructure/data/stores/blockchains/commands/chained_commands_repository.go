package commands

import (
	"path/filepath"

	stored_commands "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/commands"
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
)

// ChainedCommandsRepository represents a concrete ChainedCommandsRepository implementation
type ChainedCommandsRepository struct {
	fileRepository            stored_files.FileRepository
	cmdsRepository            stored_commands.Repository
	chainedCmdsBuilderFactory stored_commands.ChainedCommandsBuilderFactory
}

// CreateChainedCommandsRepository creates a new ChainedCommandsRepository implementation
func CreateChainedCommandsRepository(fileRepository stored_files.FileRepository, cmdsRepository stored_commands.Repository, chainedCmdsBuilderFactory stored_commands.ChainedCommandsBuilderFactory) stored_commands.ChainedCommandsRepository {
	out := ChainedCommandsRepository{
		fileRepository:            fileRepository,
		cmdsRepository:            cmdsRepository,
		chainedCmdsBuilderFactory: chainedCmdsBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a ChainedCommands instance
func (rep *ChainedCommandsRepository) Retrieve(dirPath string) (stored_commands.ChainedCommands, error) {
	metPath := filepath.Join(dirPath, "metadata.json")
	met, metErr := rep.fileRepository.Retrieve(metPath)
	if metErr != nil {
		return nil, metErr
	}

	cmdsPath := filepath.Join(dirPath, "commands")
	cmds, cmdsErr := rep.cmdsRepository.Retrieve(cmdsPath)
	if cmdsErr != nil {
		return nil, cmdsErr
	}

	prevIDPath := filepath.Join(dirPath, "previous.id")
	prevID, prevIDErr := rep.fileRepository.Retrieve(prevIDPath)
	if prevIDErr != nil {
		return nil, prevIDErr
	}

	rootIDPath := filepath.Join(dirPath, "root.id")
	rootID, rootIDErr := rep.fileRepository.Retrieve(rootIDPath)
	if rootIDErr != nil {
		return nil, rootIDErr
	}

	out, outErr := rep.chainedCmdsBuilderFactory.Create().Create().WithMetaData(met).WithCommands(cmds).WithPreviousID(prevID).WithRootID(rootID).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
