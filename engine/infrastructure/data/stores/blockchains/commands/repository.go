package commands

import (
	"path/filepath"

	stored_commands "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/commands"
	stored_chunks "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/chunks"
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
)

// Repository represents a concrete commands Repository implementation
type Repository struct {
	fileRepository     stored_files.FileRepository
	chksRepository     stored_chunks.Repository
	cmdsBuilderFactory stored_commands.BuilderFactory
}

// CreateRepository creates a new repository instance
func CreateRepository(fileRepository stored_files.FileRepository, chksRepository stored_chunks.Repository, cmdsBuilderFactory stored_commands.BuilderFactory) stored_commands.Repository {
	out := Repository{
		fileRepository:     fileRepository,
		chksRepository:     chksRepository,
		cmdsBuilderFactory: cmdsBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a stored commands
func (rep *Repository) Retrieve(dirPath string) (stored_commands.Commands, error) {
	metPath := filepath.Join(dirPath, "metadata.json")
	met, metErr := rep.fileRepository.Retrieve(metPath)
	if metErr != nil {
		return nil, metErr
	}

	//retrieve the commands:
	cmdsDirPath := filepath.Join(dirPath, "commands")
	cmds, cmdsErr := rep.chksRepository.RetrieveAll(cmdsDirPath)
	if cmdsErr != nil {
		return nil, cmdsErr
	}

	out, outErr := rep.cmdsBuilderFactory.Create().Create().WithMetaData(met).WithCommands(cmds).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
